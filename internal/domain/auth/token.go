package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *AuthController) token(c *gin.Context) {

	var input TokenInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	access_token, err := controller.service.Token(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, TokenOutput{
		AccessToken: access_token,
	})
}
