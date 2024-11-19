package email

import "time"

type SendEmailDto struct {
	TemplateSlug string `json:"template_slug" binding:"required"`
	SmtpSlug     string `json:"smtp_slug" binding:"required"`
	To           string `json:"to" binding:"required"`
	Subject      string `json:"subject"`
	Data         any    `json:"data"`
}

type EmailView struct {
	ID       uint       `json:"id"`
	SmtpName string     `json:"smtp_name"`
	From     string     `json:"from"`
	To       string     `json:"to"`
	Subject  string     `json:"subject"`
	SentAt   *time.Time `json:"sent_at"`
	HtmlBody string     `json:"html_body,omitempty"`
	TextBody string     `json:"text_body,omitempty"`
}
