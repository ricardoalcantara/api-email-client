package apikey

import (
	"github.com/ricardoalcantara/api-email-client/internal/domain"
	"github.com/ricardoalcantara/api-email-client/internal/models"
)

type ApiKeyService struct {
}

func NewApiKeyService() *ApiKeyService {
	return &ApiKeyService{}
}

func (s *ApiKeyService) post(userId uint, input CreateApiKeyDto) (*ApiKeyDto, error) {

	apiKey, fullKey, err := models.GenerateApiKey()
	if err != nil {
		return nil, err
	}
	apiKey.Name = input.Name
	apiKey.ExpiresAt = input.ExpiresAt
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
func (s *ApiKeyService) list(userId uint, pagination *models.Pagination) (*domain.ListView[ApiKeyDto], error) {
	apiKeys, err := models.ApiKeyList(userId, pagination)
	if err != nil {
		return nil, err
	}

	listView := domain.ListView[ApiKeyDto]{
		Page: pagination.Page,
	}

	for _, apiKey := range apiKeys {
		listView.List = append(listView.List, NewApiKeyDto(&apiKey))
	}

	return &listView, nil
}

func (s *ApiKeyService) get(userId, id uint) (*ApiKeyDto, error) {
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
