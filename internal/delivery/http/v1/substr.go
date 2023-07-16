package v1

import (
	"net/http"

	"github.com/begenov/tsarka-task/pkg/util"

	"github.com/gin-gonic/gin"
)

type SubstrHandler struct {
}

func NewSubstrHandler() *SubstrHandler {
	return &SubstrHandler{}
}

func (h *SubstrHandler) LoadRoutes(api *gin.RouterGroup) {

	substr := api.Group("/substr")
	{
		substr.POST("/find", h.findSubstr)
	}

}

type requestSubstr struct {
	Text string `json:"text" binding:"required"`
}

type responseSubstr struct {
	Result string `json:"result"`
}

func (s *SubstrHandler) findSubstr(ctx *gin.Context) {
	var req requestSubstr
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	str := req.Text

	res := util.MaxLengthSubstring(str)

	ctx.JSON(http.StatusOK, responseSubstr{
		Result: res,
	})
}
