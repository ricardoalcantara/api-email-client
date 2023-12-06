package email

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getSendEmail(c *gin.Context) {
	c.HTML(http.StatusOK, "pages/email/send_email.html", gin.H{})
}
