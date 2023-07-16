package v1

import (
	"hash/crc64"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HashHandler struct {
	store map[string]int
}

type Uri struct {
	Id string `uri:"id" binding:"required,min=1"`
}

type Body struct {
	Text string `json:"text" binding:"required"`
}

func NewHashHandler() *HashHandler {
	return &HashHandler{
		store: make(map[string]int),
	}
}

func (h *HashHandler) LoadRoutes(api *gin.RouterGroup) {
	hash := api.Group("/hash")
	{
		hash.POST("/calc", h.genHash)
		hash.GET("/result/:id", h.getHash)
	}
}

func (h *HashHandler) genHash(ctx *gin.Context) {
	input := Body{}
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	uid := uuid.New().String()

	go func() {
		timer := time.NewTimer(time.Minute)
		ticker := time.NewTicker(5 * time.Second)
		dec := 0

		for {
			select {
			case <-timer.C:
				log.Println("timer is done")
				h.store[uid] = dec
				return

			case <-ticker.C:
				crcTable := crc64.MakeTable(crc64.ECMA)
				checksum64 := crc64.Checksum([]byte(input.Text), crcTable)
				dec += int(time.Now().UnixNano()) & int(checksum64)
			}
		}
	}()

	h.store[uid] = 0

	ctx.JSON(http.StatusOK, gin.H{"id": uid})
}

func (h *HashHandler) getHash(ctx *gin.Context) {
	uri := Uri{}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	get, ok := h.store[uri.Id]
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is not valid"})
		return
	}

	if get == 0 {
		ctx.JSON(http.StatusOK, gin.H{"status": "PENDING"})
		return
	}

	bin := strconv.FormatInt(int64(get), 2)

	result := ""
	for _, v := range bin {
		if string(v) == "1" {
			result += string(v)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"hash": result})
}
