package email

import "time"

type EmailType string

const (
	Raw      EmailType = "raw"
	Template EmailType = "template"
	Dynamic  EmailType = "dynamic"
)

type SendEmailInput struct {
	Type         EmailType `json:"type" binding:"required"`
	TemplateName string    `json:"template_name"`
	Smtp         string    `json:"smtp"`
	To           string    `json:"to" binding:"required"`
	Subject      string    `json:"subject"`
	Context      any       `json:"context" binding:"required"`
}

type RawContext struct {
	Html string `json:"html"`
	Text string `json:"text"`
}

type EmailView struct {
	ID       uint      `json:"id"`
	SmtpName string    `json:"smtp_name"`
	From     string    `json:"from"`
	To       string    `json:"to"`
	Subject  string    `json:"subject"`
	SentAt   time.Time `json:"sent_at"`
	HtmlBody string    `json:"html_body,omitempty"`
	TextBody string    `json:"text_body,omitempty"`
}
