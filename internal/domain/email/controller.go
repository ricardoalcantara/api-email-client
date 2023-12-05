package email

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/middlewares"
)

func RegisterRoutes(r *gin.Engine) {
	routes := r.Group("/api")
	routes.Use(middlewares.AuthMiddleware())
	routes.GET("/email", List)
	routes.POST("/email", post)
	routes.GET("/email/:id", Get)
}
