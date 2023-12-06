package template

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTemplateHermes(c *gin.Context) {
	c.HTML(http.StatusOK, "pages/template/hermes.html", gin.H{})
}
