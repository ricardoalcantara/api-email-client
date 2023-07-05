package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	routes := r.Group("/api")
	routes.GET("/auth/token", Token)
}
