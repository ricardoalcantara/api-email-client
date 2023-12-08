package template

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/models"
)

func getRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "pages/template/register.html", nil)
}

func postTemplate(c *gin.Context) {
	name := c.PostForm("name")
	subject := c.PostForm("subject")
	templateHtml := c.PostForm("template-html")
	templateText := c.PostForm("template-text")

	template := models.Template{
		Name:         name,
		Subject:      subject,
		TemplateHtml: templateHtml,
		TemplateText: templateText,
	}

	if template.Save() != nil {
		c.HTML(http.StatusOK, "pages/template/register.html", gin.H{
			"name":         name,
			"subject":      subject,
			"templateHtml": templateHtml,
			"templateText": templateText,
		})
		return
	}

	c.Redirect(http.StatusFound, "/template")
}
