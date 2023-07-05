package email

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/emailengine"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/internal/utils"
)

func SendEmail(c *gin.Context) {
	var input SendEmailInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": utils.PrintError(err)})
		return
	}

	var smtp *models.Smtp
	var err error
	if len(input.Smtp) == 0 {
		smtp, err = models.SmtpGetDefault()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": utils.PrintError(err)})
			return
		}
	} else {
		smtp, err = models.SmtpGetByName(input.Smtp)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": utils.PrintError(err)})
			return
		}
	}

	var html, text, subject string
	switch input.Type {
	case Raw:
		data, err := utils.TypeConverter[RawContext](&input.Context)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": utils.PrintError(err)})
			return
		}

		if len(data.Html) == 0 && len(data.Text) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Html or Text are required"})
			return
		}

		html = data.Html
		text = data.Text
		subject = input.Subject
	case Dynamic:
		html, text = emailengine.GetHermes(input.Context)

		if len(input.Subject) > 0 {
			subject = input.Subject
		} else {
			subject = "GetHermes"
		}
	case Template:
		t, err := models.TemplateGet("Default")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": utils.PrintError(err)})
			return
		}

		if len(t.TemplateHtml) > 0 {
			html = emailengine.GetTemplate(t.TemplateHtml, input.Context)
		} else {
			html = ""
		}

		if len(t.TemplateText) > 0 {
			text = emailengine.GetTemplate(t.TemplateText, input.Context)
		} else {
			text = ""
		}

		if len(input.Subject) > 0 {
			subject = input.Subject
		} else {
			subject = t.Subject
		}

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

	// go func() {
	// 	if err = emailengine.SendEmail(dialer, smtp.Email, input.To, subject, html, text); err != nil {
	// 		utils.PrintError(err)
	// 	}
	// 	logrus.Info("Email sent!")
	// }()

	c.JSON(http.StatusAccepted, gin.H{})
}
