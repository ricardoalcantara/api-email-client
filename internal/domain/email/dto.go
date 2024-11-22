package email

import (
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/pkg/types"
)

func NewEmailDto(email *models.Email) types.EmailDto {
	var smtpName string
	var from string
	if email.Smtp != nil {
		smtpName = email.Smtp.Name
		from = email.Smtp.Email
	}
	return types.EmailDto{
		ID:       email.ID,
		SmtpName: smtpName,
		From:     from,
		To:       email.To,
		Subject:  email.Subject,
		SentAt:   email.SentAt,
		HtmlBody: email.HtmlBody,
		TextBody: email.TextBody,
	}
}
