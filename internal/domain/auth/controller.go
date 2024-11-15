package auth

import (
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service AuthService
}

func RegisterRoutes(r *gin.Engine) {
	controller := AuthController{}
	routes := r.Group("/api")
	routes.POST("/auth/token", controller.token)
}
