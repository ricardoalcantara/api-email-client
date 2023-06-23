package email

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/emailengine"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/internal/utils"
)

type Person struct {
	Name string
}

type EmailType string

const (
	Raw      EmailType = "raw"
	Template EmailType = "template"
	Dynamic  EmailType = "dynamic"
)

type SendEmailInput struct {
	Type         EmailType `json:"type" binding:"required"`
	TemplateName string    `json:"template_name"`
	Smtp         string    `json:"smtp" binding:"required"`
	To           string    `json:"to" binding:"required"`
	Subject      string    `json:"subject"`
	Context      any       `json:"context" binding:"required"`
}

type RawContext struct {
	Html string `json:"html"`
	Text string `json:"text"`
}

func SendEmail(c *gin.Context) {
	var input SendEmailInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": utils.PrintError(err)})
		return
	}

	smtp, err := models.SmtpGet(input.Smtp)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": utils.PrintError(err)})
		return
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
			html = emailengine.GetTemplate(t.TemplateHtml, Person{Name: "Ricardo"})
		} else {
			html = ""
		}

		if len(t.TemplateText) > 0 {
			text = emailengine.GetTemplate(t.TemplateText, Person{Name: "Ricardo"})
		} else {
			text = ""
		}

		if len(input.Subject) > 0 {
			subject = input.Subject
		} else {
			subject = t.Subject
		}

	}

	dialer := smtp.GetDialer()

	if err = emailengine.SendEmail(dialer, smtp.Email, input.To, subject, html, text); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": utils.PrintErrorAnd(err, "Could not connect to SMTP server")})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{})
}
