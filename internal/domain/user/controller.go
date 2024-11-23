package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/middlewares"
	"github.com/ricardoalcantara/api-email-client/pkg/types"
)

type UserController struct {
	service *UserService
}

func NewUserController() *UserController {
	return &UserController{
		service: NewUserService(),
	}
}

func (controller *UserController) password(c *gin.Context) {
	var input types.UpdatePasswordDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := strconv.Atoi(c.GetString("x-id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.service.UpdatePassword(uint(userId), input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusAccepted)
}

func RegisterRoutes(r *gin.Engine) {
	controller := NewUserController()

	routes := r.Group("/api")
	routes.Use(middlewares.AuthMiddleware())
	routes.PATCH("/user/password", controller.password)
}
