package email

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/middlewares"
)

func RegisterRoutes(r *gin.Engine) {
	routes := r.Group("/api")
	routes.Use(middlewares.AuthMiddleware(os.Getenv("JWT_SECRET")))
	routes.POST("/email", SendEmail)
}
