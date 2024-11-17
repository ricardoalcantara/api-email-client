package template

import (
	"github.com/ricardoalcantara/api-email-client/internal/domain"
	"github.com/ricardoalcantara/api-email-client/internal/models"
)

type TemplateService struct {
}

func NewTemplateService() *TemplateService {
	return &TemplateService{}
}

func (s *TemplateService) Create(templateDto *CreateTemplateDto) (*TemplateDto, error) {
	template := models.Template{
		Name:         templateDto.Name,
		Slug:         templateDto.Slug,
		JsonSchema:   templateDto.JsonSchema,
		Subject:      templateDto.Subject,
		TemplateHtml: templateDto.TemplateHtml,
		TemplateText: templateDto.TemplateText,
	}
	err := template.Save()
	if err != nil {
		return nil, err
	}

	return NewTemplateDto(&template), nil
}

func (s *TemplateService) Get(id uint) (*TemplateDto, error) {
	template, err := models.TemplateGet(id)
	if err != nil {
		return nil, err
	}

	return NewTemplateDto(template), nil
}

func (s *TemplateService) List(pagination *models.Pagination) (*domain.ListView[TemplateDto], error) {
	templates, err := models.TemplateList(pagination)
	if err != nil {
		return nil, err
	}

	listView := domain.ListView[TemplateDto]{}
	for _, t := range templates {
		listView.List = append(listView.List, TemplateDto{
			ID:           t.ID,
			Name:         t.Name,
			Subject:      t.Subject,
			TemplateHtml: t.TemplateHtml,
			TemplateText: t.TemplateText,
		})
	}

	return &listView, nil
}

func (s *TemplateService) Update(id uint, updateTemplate *UpdateTemplateDto) (*TemplateDto, error) {
	template, err := models.TemplateGet(id)
	if err != nil {
		return nil, err
	}

	update := map[string]interface{}{}
	if updateTemplate.Name != nil {
		update["name"] = *updateTemplate.Name
	}
	if updateTemplate.Slug != nil {
		update["slug"] = *updateTemplate.Slug
	}
	if updateTemplate.JsonSchema != nil {
		update["json_schema"] = *updateTemplate.JsonSchema
	}
	if updateTemplate.Subject != nil {
		update["subject"] = *updateTemplate.Subject
	}
	if updateTemplate.TemplateHtml != nil {
		update["template_html"] = *updateTemplate.TemplateHtml
	}
	if updateTemplate.TemplateText != nil {
		update["template_text"] = *updateTemplate.TemplateText
	}

	if len(update) > 0 {
		err = template.Updates(update)
		if err != nil {
			return nil, err
		}
	}
	return NewTemplateDto(template), nil
}

func (s *TemplateService) Delete(id uint) error {
	return models.TemplateDelete(id)
}
