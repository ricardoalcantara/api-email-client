package email

import (
	"github.com/ricardoalcantara/api-email-client/internal/domain"
	"github.com/ricardoalcantara/api-email-client/internal/emailengine"
	"github.com/ricardoalcantara/api-email-client/internal/models"
)

type EmailService struct {
}

func NewEmailService() *EmailService {
	return &EmailService{}
}

func (s *EmailService) post(input SendEmailDto) error {
	var smtp *models.Smtp
	var err error
	if len(input.SmtpSlug) == 0 {
		smtp, err = models.SmtpGetDefault()
		if err != nil {
			return err
		}
	} else {
		smtp, err = models.SmtpGetBySlug(input.SmtpSlug)
		if err != nil {

			return err
		}
	}

	t, err := models.TemplateGetBySlug(input.TemplateSlug)
	if err != nil {

		return err
	}

	html := emailengine.GetTemplate(t.TemplateHtml, input.Data)
	text := emailengine.GetTemplate(t.TemplateText, input.Data)

	var subject string
	if len(input.Subject) > 0 {
		subject = input.Subject
	} else {
		subject = t.Subject
	}

	email := models.Email{
		SmtpId:   smtp.ID,
		To:       input.To,
		Subject:  subject,
		HtmlBody: html,
		TextBody: text,
	}

	if err := email.Save(); err != nil {
		return err
	}

	email.Smtp = smtp

	if err = emailengine.SendEmailQueue(email); err != nil {
		return err
	}

	return nil
}

func (s *EmailService) list(pagination *models.Pagination) (*domain.ListView[EmailView], error) {
	emails, err := models.EmailList(pagination)
	if err != nil {
		return nil, err
	}

	result := domain.ListView[EmailView]{Page: pagination.Page}

	for _, e := range emails {
		result.List = append(result.List, NewEmailView(&e))
	}

	return &result, nil
}

func (s *EmailService) get(id uint) (*EmailView, error) {
	email, err := models.EmailGet(uint(id))
	if err != nil {
		return nil, err
	}

	emailView := NewEmailView(email)

	return &emailView, nil
}
