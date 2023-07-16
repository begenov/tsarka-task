package v1

import (
	"net/http"

	"github.com/begenov/tsarka-task/internal/service"

	"github.com/gin-gonic/gin"
)

type CounterHandler struct {
	service service.Couters
}

func NewCounterHandler(service service.Couters) *CounterHandler {
	return &CounterHandler{
		service: service,
	}
}

func (h *CounterHandler) LoadRoutes(api *gin.RouterGroup) {
	counter := api.Group("counter")
	{
		counter.POST("/add/:i", h.addCounter)
		counter.POST("/sub/:i", h.subCounter)
		counter.GET("/val", h.valCounter)
	}
}

const (
	key = "counter"
)

type counterRequest struct {
	I int64 `uri:"i" binding:"required,min=1"`
}

type counterResponse struct {
	Counter int64 `json:"counter"`
}

func (h *CounterHandler) addCounter(ctx *gin.Context) {
	var req counterRequest
	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	res, err := h.service.Add(key, req.I)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Can't to add value"})
		return
	}

	ctx.JSON(http.StatusOK, counterResponse{Counter: res})
}

func (h *CounterHandler) subCounter(ctx *gin.Context) {

	var req counterRequest
	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	res, err := h.service.Sub(key, req.I)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Can't to add sub"})
		return
	}

	ctx.JSON(http.StatusOK, counterResponse{Counter: res})
}

func (h *CounterHandler) valCounter(ctx *gin.Context) {
	res, err := h.service.Get(key)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Can't to get value"})
		return
	}

	ctx.JSON(http.StatusOK, counterResponse{Counter: res})
}
