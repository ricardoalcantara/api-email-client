package client

import (
	"fmt"
	"net/http"

	"github.com/ricardoalcantara/api-email-client/pkg/types"
)

type ApiClient struct {
	baseURL             string
	httpClient          *http.Client
	authorizetionHeader string
}

type Option func(*ApiClient)

func WithBaseURL(baseURL string) Option {
	return func(c *ApiClient) {
		c.baseURL = baseURL
	}
}

func WithToken(token string) Option {
	return func(c *ApiClient) {
		c.authorizetionHeader = "Bearer " + token
	}
}

func WithApiKey(apiKey string) Option {
	return func(c *ApiClient) {
		c.authorizetionHeader = "ApiKey " + apiKey
	}
}

func New(opts ...Option) *ApiClient {
	c := &ApiClient{
		baseURL:    "http://localhost:8080",
		httpClient: &http.Client{},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// Auth methods
func (c *ApiClient) GetToken(email, password string) (*ApiResponse[types.TokenOutput], error) {
	return doRequest[types.TokenOutput](c, http.MethodPost, "/api/auth/token", types.TokenInput{
		Email:    email,
		Password: password,
	})
}

// Email methods
func (c *ApiClient) ListEmails(page int) (*ApiResponse[types.ListView[types.EmailDto]], error) {
	return doRequest[types.ListView[types.EmailDto]](c, http.MethodGet, fmt.Sprintf("/api/email?page=%d", page), nil)
}

func (c *ApiClient) GetEmail(id uint) (*ApiResponse[types.EmailDto], error) {
	return doRequest[types.EmailDto](c, http.MethodGet, fmt.Sprintf("/api/email/%d", id), nil)
}

func (c *ApiClient) SendEmail(req *types.SendEmailDto) (*ApiResponse[struct{}], error) {
	return doRequest[struct{}](c, http.MethodPost, "/api/email", req)
}

// SMTP methods
func (c *ApiClient) ListSMTPConfigs(page int) (*ApiResponse[types.ListView[types.SmtpDto]], error) {
	return doRequest[types.ListView[types.SmtpDto]](c, http.MethodGet, fmt.Sprintf("/api/smtp?page=%d", page), nil)
}

func (c *ApiClient) GetSMTPConfig(slug string) (*ApiResponse[types.SmtpDto], error) {
	return doRequest[types.SmtpDto](c, http.MethodGet, fmt.Sprintf("/api/smtp/%s", slug), nil)
}

func (c *ApiClient) CreateSMTPConfig(config *types.CreateSmtpDto) (*ApiResponse[struct{}], error) {
	return doRequest[struct{}](c, http.MethodPost, "/api/smtp", config)
}

func (c *ApiClient) UpdateSMTPConfig(slug string, config *types.UpdateSmtpDto) (*ApiResponse[struct{}], error) {
	return doRequest[struct{}](c, http.MethodPut, fmt.Sprintf("/api/smtp/%s", slug), config)
}

func (c *ApiClient) PatchSMTPConfig(slug string, updates map[string]interface{}) (*ApiResponse[struct{}], error) {
	return doRequest[struct{}](c, http.MethodPatch, fmt.Sprintf("/api/smtp/%s", slug), updates)
}

func (c *ApiClient) DeleteSMTPConfig(slug string) (*ApiResponse[struct{}], error) {
	return doRequest[struct{}](c, http.MethodDelete, fmt.Sprintf("/api/smtp/%s", slug), nil)
}

// Template methods
func (c *ApiClient) ListTemplates(page int) (*ApiResponse[types.ListView[types.TemplateDto]], error) {
	return doRequest[types.ListView[types.TemplateDto]](c, http.MethodGet, fmt.Sprintf("/api/template?page=%d", page), nil)
}

func (c *ApiClient) GetTemplate(slug string) (*ApiResponse[types.TemplateDto], error) {
	return doRequest[types.TemplateDto](c, http.MethodGet, fmt.Sprintf("/api/template/%s", slug), nil)
}

func (c *ApiClient) CreateTemplate(template *types.CreateTemplateDto) (*ApiResponse[struct{}], error) {
	return doRequest[struct{}](c, http.MethodPost, "/api/template", template)
}

func (c *ApiClient) UpdateTemplate(slug string, template *types.UpdateTemplateDto) (*ApiResponse[struct{}], error) {
	return doRequest[struct{}](c, http.MethodPut, fmt.Sprintf("/api/template/%s", slug), template)
}

func (c *ApiClient) PatchTemplate(slug string, updates map[string]interface{}) (*ApiResponse[struct{}], error) {
	return doRequest[struct{}](c, http.MethodPatch, fmt.Sprintf("/api/template/%s", slug), updates)
}

func (c *ApiClient) DeleteTemplate(slug string) (*ApiResponse[struct{}], error) {
	return doRequest[struct{}](c, http.MethodDelete, fmt.Sprintf("/api/template/%s", slug), nil)
}

// API Key methods
func (c *ApiClient) ListAPIKeys(page int) (*ApiResponse[types.ListView[types.ApiKeyDto]], error) {
	return doRequest[types.ListView[types.ApiKeyDto]](c, http.MethodGet, fmt.Sprintf("/api/api-key?page=%d", page), nil)
}

func (c *ApiClient) GetAPIKey(id uint) (*ApiResponse[types.ApiKeyDto], error) {
	return doRequest[types.ApiKeyDto](c, http.MethodGet, fmt.Sprintf("/api/api-key/%d", id), nil)
}

func (c *ApiClient) CreateAPIKey(key *types.CreateApiKeyDto) (*ApiResponse[types.ApiKeyDto], error) {
	return doRequest[types.ApiKeyDto](c, http.MethodPost, "/api/api-key", key)
}

func (c *ApiClient) RegenerateAPIKey(id uint) (*ApiResponse[struct{}], error) {
	return doRequest[struct{}](c, http.MethodPatch, fmt.Sprintf("/api/api-key/%d/regenerate", id), nil)
}

func (c *ApiClient) DeleteAPIKey(id uint) (*ApiResponse[struct{}], error) {
	return doRequest[struct{}](c, http.MethodDelete, fmt.Sprintf("/api/api-key/%d", id), nil)
}
