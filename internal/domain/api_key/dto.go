package apikey

import (
	"time"

	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/pkg/types"
)

func NewApiKeyDto(a *models.ApiKey) types.ApiKeyDto {
	dtp := types.ApiKeyDto{
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
