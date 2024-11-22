package types

import (
	"time"
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
