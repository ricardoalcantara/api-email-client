package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHome(c *gin.Context) {
	c.HTML(http.StatusOK, "pages/home/index.html", nil)
}
