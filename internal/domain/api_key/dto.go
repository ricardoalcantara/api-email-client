package apikey

import (
	"time"

	"github.com/ricardoalcantara/api-email-client/internal/models"
)

type CreateApiKeyDto struct {
	Name        string `json:"name" binding:"required"`
	IpWhitelist string `json:"ip_whitelist"`
	ExpiresAt   string `json:"expires_at"`
}

type ApiKeyDto struct {
	Id          uint       `json:"id"`
	Name        string     `json:"name"`
	Key         string     `json:"key,omitempty"`
	LastUsed    *time.Time `json:"last_used"`
	IpWhitelist string     `json:"ip_whitelist"`
	ExpiresAt   string     `json:"expires_at"`
}

func NewApiKeyDto(a *models.ApiKey) ApiKeyDto {
	dtp := ApiKeyDto{
		Id:          a.ID,
		Name:        a.Name,
		LastUsed:    a.LastUsed,
		IpWhitelist: a.IpWhitelist,
	}

	if a.ExpiresAt != nil {
		dtp.ExpiresAt = a.ExpiresAt.Format(time.RFC3339)
	}

	return dtp
}
