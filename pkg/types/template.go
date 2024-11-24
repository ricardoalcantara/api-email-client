package types

import "github.com/matcornic/hermes/v2"

type TemplateDto struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	JsonSchema   string `json:"json_schema"`
	Subject      string `json:"subject"`
	TemplateHtml string `json:"template_html"`
	TemplateText string `json:"template_text"`
}

type CreateTemplateDto struct {
	Name         string `json:"name" binding:"required"`
	Slug         string `json:"slug" binding:"required"`
	JsonSchema   string `json:"json_schema" `
	Subject      string `json:"subject" binding:"required"`
	TemplateHtml string `json:"template_html"`
	TemplateText string `json:"template_text"`
}

type UpdateTemplateDto struct {
	Name         *string `json:"name"`
	Slug         *string `json:"slug"`
	JsonSchema   *string `json:"json_schema"`
	Subject      *string `json:"subject"`
	TemplateHtml *string `json:"template_html"`
	TemplateText *string `json:"template_text"`
}

type RequestTemplateGeneratorDto struct {
	Theme string          `json:"theme"`
	Config hermes.Hermes `json:"config"`
	Email  hermes.Body   `json:"email"`
}

type TemplateGeneratorDto struct {
	TemplateHtml string `json:"template_html"`
	TemplateText string `json:"template_text"`
}
