package v1

import (
	"net/http"

	"github.com/begenov/tsarka-task/pkg/util"

	"github.com/gin-gonic/gin"
)

type InnEmailHandler struct {
}

func NewEmailInnHandler() *InnEmailHandler {
	return &InnEmailHandler{}
}

func (h *InnEmailHandler) LoadRoutes(api *gin.RouterGroup) {
	check := api.Group("/check")
	{
		check.POST("/email", h.checkEmails)
		check.POST("/inn", h.checkInn)
	}
}

type emailRequest struct {
	Emails []string `json:"emails" binding:"required"`
}

type responseEmails struct {
	Emails []string `json:"valid_emails"`
}

func (h *InnEmailHandler) checkEmails(ctx *gin.Context) {
	var req emailRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	emails := util.EmailsCheck(req.Emails)

	ctx.JSON(http.StatusOK, responseEmails{Emails: emails})
}

type innRequest struct {
	Inn []string `json:"inn" binding:"required"`
}

type innResponse struct {
	Inn []string `json:"inn"`
}

func (h *InnEmailHandler) checkInn(ctx *gin.Context) {
	var req innRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	inn := util.InnCheck(req.Inn)

	ctx.JSON(http.StatusOK, innResponse{
		Inn: inn,
	})
}
