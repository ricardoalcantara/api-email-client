package template

import (
	"github.com/matcornic/hermes/v2"
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

func NewTemplateGeneratorDto(t *types.RequestTemplateGeneratorDto) (*types.TemplateGeneratorDto, error) {
	templateGeneratorDto := types.TemplateGeneratorDto{}

	email := hermes.Email{
		Body: t.Email,
	}
	var err error
	templateGeneratorDto.TemplateHtml, err = t.Config.GenerateHTML(email)
	if err != nil {
		return nil, err
	}

	templateGeneratorDto.TemplateText, err = t.Config.GeneratePlainText(email)
	if err != nil {
		return nil, err
	}
	return &templateGeneratorDto, nil
}
