package template

import (
	"github.com/gin-gonic/gin"
)

func Post(c *gin.Context) {
	// h.DB
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
