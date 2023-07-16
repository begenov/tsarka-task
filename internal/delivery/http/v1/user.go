package v1

import (
	"net/http"

	"github.com/begenov/tsarka-task/internal/domain"
	"github.com/begenov/tsarka-task/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.Users
}

func NewUserHandler(service service.Users) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) LoadRoutes(api *gin.RouterGroup) {
	user := api.Group("/user")
	{
		user.POST("", h.createUser)
		user.GET("/:id", h.getUser)
		user.PUT("/:id", h.updateUser)
		user.DELETE("/:id", h.deleteUser)
	}
}

type uriUser struct {
	Id int `uri:"id" binding:"required,min=1"`
}

func (h *UserHandler) createUser(ctx *gin.Context) {
	var req domain.User
	var err error
	if err = ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	req.ID, err = h.service.CreateUser(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Can't to create user"})
		return
	}

	ctx.JSON(http.StatusOK, req)
}

func (h *UserHandler) getUser(ctx *gin.Context) {
	req := uriUser{}
	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := h.service.GetUser(ctx, req.Id)
	if err != nil {
		if err == domain.NotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Not found user"})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Can't to get user"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) updateUser(ctx *gin.Context) {
	var req domain.User
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	inp := uriUser{}
	if err := ctx.BindUri(&inp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Uri"})
		return
	}
	req.ID = inp.Id
	user, err := h.service.UpdateUser(ctx, req)
	if err != nil {
		if err == domain.NotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Not found user"})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Can't to update user"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) deleteUser(ctx *gin.Context) {

	inp := uriUser{}
	if err := ctx.BindUri(&inp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Uri"})
		return
	}
	err := h.service.DeleteUser(ctx, inp.Id)
	if err != nil {
		if err == domain.NotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Not found user"})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Can't to delete user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Status OK"})
}
