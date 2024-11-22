package types

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
