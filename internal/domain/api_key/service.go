package apikey

import (
	"time"

	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/pkg/types"
)

type ApiKeyService struct {
}

func NewApiKeyService() *ApiKeyService {
	return &ApiKeyService{}
}

func (s *ApiKeyService) post(userId uint, input types.CreateApiKeyDto) (*types.ApiKeyDto, error) {

	apiKey, fullKey, err := models.GenerateApiKey()
	if err != nil {
		return nil, err
	}
	apiKey.Name = input.Name
	if len(input.ExpiresAt) > 0 {
		expiresAt, err := time.Parse(time.DateOnly, input.ExpiresAt)
		if err != nil {
			return nil, err
		}
		apiKey.ExpiresAt = &expiresAt
	}
	apiKey.IpWhitelist = input.IpWhitelist
	apiKey.UserId = userId

	err = apiKey.Save()
	if err != nil {
		return nil, err
	}

	view := NewApiKeyDto(apiKey)
	view.Key = fullKey

	return &view, nil
}
func (s *ApiKeyService) list(userId uint, pagination *models.Pagination) (*types.ListView[types.ApiKeyDto], error) {
	apiKeys, err := models.ApiKeyList(userId, pagination)
	if err != nil {
		return nil, err
	}

	listView := types.ListView[types.ApiKeyDto]{
		Page: pagination.Page,
	}

	for _, apiKey := range apiKeys {
		listView.List = append(listView.List, NewApiKeyDto(&apiKey))
	}

	return &listView, nil
}

func (s *ApiKeyService) get(userId, id uint) (*types.ApiKeyDto, error) {
	apiKey, err := models.ApiKeyGet(userId, id)
	if err != nil {
		return nil, err
	}

	result := NewApiKeyDto(apiKey)

	return &result, nil
}

func (s *ApiKeyService) delete(userId, id uint) error {
	return models.ApiKeyDelete(userId, id)
}
