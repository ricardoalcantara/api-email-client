package template

import (
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/pkg/types"
)

func NewTemplateDto(t *models.Template) types.TemplateDto {
	return types.TemplateDto{
		ID:           t.ID,
		Name:         t.Name,
		Slug:         t.Slug,
		JsonSchema:   t.JsonSchema,
		Subject:      t.Subject,
		TemplateHtml: t.TemplateHtml,
		TemplateText: t.TemplateText,
	}
}
