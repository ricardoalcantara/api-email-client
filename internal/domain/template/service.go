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

	view := NewTemplateDto(&template)
	return &view, nil
}

func (s *TemplateService) Get(slug string) (*TemplateDto, error) {
	template, err := models.TemplateGetBySlug(slug)
	if err != nil {
		return nil, err
	}

	view := NewTemplateDto(template)
	return &view, nil
}

func (s *TemplateService) List(pagination *models.Pagination) (*domain.ListView[TemplateDto], error) {
	templates, err := models.TemplateList(pagination)
	if err != nil {
		return nil, err
	}

	listView := domain.ListView[TemplateDto]{}
	for _, t := range templates {
		listView.List = append(listView.List, NewTemplateDto(&t))
	}

	return &listView, nil
}

func (s *TemplateService) Patch(slug string, updateTemplate *UpdateTemplateDto) (*TemplateDto, error) {
	template, err := models.TemplateGetBySlug(slug)
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
	view := NewTemplateDto(template)
	return &view, nil
}

func (s *TemplateService) Update(slug string, updateTemplate *UpdateTemplateDto) (*TemplateDto, error) {
	template, err := models.TemplateGetBySlug(slug)
	if err != nil {
		return nil, err
	}

	template.Name = *updateTemplate.Name
	template.Slug = *updateTemplate.Slug
	template.JsonSchema = *updateTemplate.JsonSchema
	template.Subject = *updateTemplate.Subject
	template.TemplateHtml = *updateTemplate.TemplateHtml
	template.TemplateText = *updateTemplate.TemplateText

	err = template.Update()
	if err != nil {
		return nil, err
	}

	view := NewTemplateDto(template)
	return &view, nil
}

func (s *TemplateService) Delete(slug string) error {
	template, err := models.TemplateGetBySlug(slug)
	if err != nil {
		return err
	}
	return template.Delete()
}
