package email

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/emailengine"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/internal/utils"
)

func post(c *gin.Context) {
	var input SendEmailDto

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": utils.PrintError(err)})
		return
	}

	var smtp *models.Smtp
	var err error
	if len(input.SmtpSlug) == 0 {
		smtp, err = models.SmtpGetDefault()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": utils.PrintError(err)})
			return
		}
	} else {
		smtp, err = models.SmtpGetBySlug(input.SmtpSlug)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": utils.PrintError(err)})
			return
		}
	}

	t, err := models.TemplateGetBySlug(input.TemplateSlug)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": utils.PrintError(err)})
		return
	}

	html := emailengine.GetTemplate(t.TemplateHtml, input.Data)
	text := emailengine.GetTemplate(t.TemplateText, input.Data)

	var subject string
	if len(input.Subject) > 0 {
		subject = input.Subject
	} else {
		subject = t.Subject
	}

	email := models.Email{
		SmtpId:   smtp.ID,
		To:       input.To,
		Subject:  subject,
		HtmlBody: html,
		TextBody: text,
	}

	if err := email.Save(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": utils.PrintErrorAnd(err, "Fail sending email")})
		return
	}

	email.Smtp = smtp

	if err = emailengine.SendEmailQueue(email); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": utils.PrintErrorAnd(err, "Fail sending email")})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{})
}
