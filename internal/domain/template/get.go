package template

import (
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	// h.DB
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
