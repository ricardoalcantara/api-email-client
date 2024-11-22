package smtp

import (
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/pkg/types"
)

func NewSmtpDto(s *models.Smtp) types.SmtpDto {
	return types.SmtpDto{
		ID:      s.ID,
		Name:    s.Name,
		Slug:    s.Slug,
		Server:  s.Server,
		Port:    s.Port,
		Email:   s.Email,
		User:    s.User,
		Default: s.Default,
	}
}
