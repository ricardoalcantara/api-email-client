package smtp

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/middlewares"
)

func RegisterRoutes(r *gin.Engine) {

	authorized := r.Group("/")
	authorized.Use(middlewares.SessionAuthentication())
	authorized.GET("/smtp", getIndex)

	authorized.GET("/smtp/register", getRegister)
	authorized.POST("/smtp/register", postSmtp)

	authorized.DELETE("/smtp/delete", deleteSmtp)
}
