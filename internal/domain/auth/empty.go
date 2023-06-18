package auth

import (
	"github.com/gin-gonic/gin"
)

func Empty(c *gin.Context) {
	// h.DB
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
