package email

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/emailengine"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/internal/utils"
)

func post(c *gin.Context) {
	var input SendEmailInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": utils.PrintError(err)})
		return
	}

	var smtp *models.Smtp
	var err error
	if input.SmtpId == 0 {
		smtp, err = models.SmtpGetDefault()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": utils.PrintError(err)})
			return
		}
	} else {
		smtp, err = models.SmtpGetById(input.SmtpId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": utils.PrintError(err)})
			return
		}
	}

	t, err := models.TemplateGet(input.TemplateId)
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

	// dialer, err := smtp.GetDialer()
	// *smtp, input.To, subject, html, text)
	if err = emailengine.SendEmailQueue(email); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": utils.PrintErrorAnd(err, "Fail sending email")})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{})
}
