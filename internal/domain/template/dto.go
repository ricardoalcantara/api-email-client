package template

type TemplateView struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	JsonSchema   string `json:"json_schema"`
	Subject      string `json:"subject"`
	TemplateHtml string `json:"template_html"`
	TemplateText string `json:"template_text"`
}
