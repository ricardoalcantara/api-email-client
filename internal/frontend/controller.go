package frontend

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/middlewares"
)

func RegisterRoutes(r *gin.Engine) {

	r.GET("/login", getLogin)
	r.POST("/login", postLogin)
	r.GET("/logout", getLogout)
	r.GET("/about", getAbout)
	r.NoRoute(notFound)

	authorized := r.Group("/")
	authorized.Use(middlewares.SessionAuthentication())
	authorized.GET("/", getHome)
	authorized.GET("/smtp", getSmtp)
	authorized.POST("/smtp", postSmtp)
	authorized.POST("/smtp/delete", deleteSmtp)
	authorized.GET("/template", getTemplate)
	authorized.GET("/email", getEmail)
}

func notFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "pages/404.html", nil)
}
