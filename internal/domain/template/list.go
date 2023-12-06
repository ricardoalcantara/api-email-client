package template

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/api-email-client/internal/domain"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/internal/utils"
	"github.com/samber/lo"
)

func List(c *gin.Context) {
	p := models.NewPagination(c)
	templates, err := models.TemplateList(p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	result := lo.Map(templates, func(t models.Template, index int) TemplateView {
		return TemplateView{
			ID:           t.ID,
			Name:         t.Name,
			Subject:      t.Subject,
			TemplateHtml: t.TemplateHtml,
			TemplateText: t.TemplateText,
		}
	})

	c.JSON(http.StatusOK, domain.ListView[TemplateView]{List: result, Page: p.Page})
}
