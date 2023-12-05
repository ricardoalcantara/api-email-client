package smtp

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/middlewares"
)

func RegisterRoutes(r *gin.Engine) {

	authorized := r.Group("/")
	authorized.Use(middlewares.SessionAuthentication())
	authorized.GET("/smtp", getSmtp)
	authorized.POST("/smtp", postSmtp)
	authorized.POST("/smtp/delete", deleteSmtp)
}
