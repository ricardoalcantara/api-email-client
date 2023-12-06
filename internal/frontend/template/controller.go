package template

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/middlewares"
)

func RegisterRoutes(r *gin.Engine) {

	authorized := r.Group("/")
	authorized.Use(middlewares.SessionAuthentication())
	authorized.GET("/template", getIndex)

	authorized.GET("/template/register", getRegister)
	authorized.POST("/template/register", postTemplate)

	authorized.GET("/template/hermes", getTemplateHermes)

	// authorized.DELETE("/template/delete", deleteTemplate)
}
