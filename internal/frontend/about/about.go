package about

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAbout(c *gin.Context) {
	c.HTML(http.StatusOK, "pages/about/index.html", nil)
}
