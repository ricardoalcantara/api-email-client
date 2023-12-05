package testemail

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTestEmail(c *gin.Context) {
	c.HTML(http.StatusOK, "pages/test_email/index.html", gin.H{})
}
