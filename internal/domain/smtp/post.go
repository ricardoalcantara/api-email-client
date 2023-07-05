package smtp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/domain"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/internal/utils"
)

func Post(c *gin.Context) {
	var input SmtpRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	if input.Default {
		if err := models.SmtpDisableDefault(); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
			return
		}
	}

	smtp := models.Smtp{
		Name:     input.Name,
		Server:   input.Server,
		Port:     input.Port,
		Email:    input.Email,
		User:     input.User,
		Password: input.Password,
		Default:  input.Default,
	}

	smtp.Base64Password()
	err := smtp.Save()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	c.JSON(http.StatusAccepted, SmtpView{
		ID:      smtp.ID,
		Name:    smtp.Name,
		Server:  smtp.Server,
		Port:    smtp.Port,
		Email:   smtp.Email,
		User:    smtp.User,
		Default: smtp.Default,
	})
}
