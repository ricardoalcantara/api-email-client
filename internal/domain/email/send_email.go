package email

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/emailengine"
	"github.com/ricardoalcantara/api-email-client/internal/models"
)

type Person struct {
	Name string
}

type SendEmailInput struct {
	Type         string `json:"type" binding:"required"`
	TemplateName string `json:"template_name" binding:"required"`
	SmtpId       uint   `json:"smtp_id" binding:"required"`
	To           string `json:"to" binding:"required"`
	Subject      string `json:"subject"`
	Context      any    `json:"context" binding:"required"`
}

func SendEmail(c *gin.Context) {
	var input SendEmailInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	html, text := emailengine.GetHermes()

	var subject string

	if len(input.Subject) > 0 {
		subject = input.Subject
	} else {
		subject = "GetHermes"
	}
	emailengine.SendEmail("admin@admin.com", input.To, subject, html, text)

	t, err := models.TemplateGet("Default")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	emailengine.SendEmail("admin@admin.com", input.To, subject, html, text)

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
