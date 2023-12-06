package template

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/domain"
	"github.com/ricardoalcantara/api-email-client/internal/domain/template"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/internal/utils"
	"github.com/samber/lo"
)

func getIndex(c *gin.Context) {
	p := models.NewPagination(c)
	templates, err := models.TemplateList(p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	result := lo.Map(templates, func(t models.Template, index int) template.TemplateView {
		return template.TemplateView{
			ID:           t.ID,
			Name:         t.Name,
			JsonSchema:   t.JsonSchema,
			Subject:      t.Subject,
			TemplateHtml: t.TemplateHtml,
			TemplateText: t.TemplateText,
		}
	})

	c.HTML(http.StatusOK, "pages/template/index.html", gin.H{
		"listTemplate": result,
	})
}
