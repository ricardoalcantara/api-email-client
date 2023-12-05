package template

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/middlewares"
)

func RegisterRoutes(r *gin.Engine) {

	authorized := r.Group("/")
	authorized.Use(middlewares.SessionAuthentication())
	authorized.GET("/template", getTemplate)
}

func notFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "pages/404.html", nil)
}
