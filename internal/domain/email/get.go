package email

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/domain"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/internal/utils"
)

func Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	email, err := models.EmailGet(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	c.JSON(http.StatusOK, EmailView{
		ID:       email.ID,
		SmtpName: email.Smtp.Name,
		From:     email.Smtp.Email,
		To:       email.To,
		Subject:  email.Subject,
		SentAt:   email.SentAt,
		HtmlBody: email.HtmlBody,
		TextBody: email.TextBody,
	})
}
