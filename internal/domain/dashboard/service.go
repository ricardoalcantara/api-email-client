package dashboard

import (
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/pkg/types"
)

type DashboardService struct {
}

func NewDashboardService() *DashboardService {
	return &DashboardService{}
}

func (s *DashboardService) Get() (*types.DashboardDto, error) {
	templates, err := models.TemplateCount()
	if err != nil {
		return nil, err
	}
	emails, err := models.EmailCount()
	if err != nil {
		return nil, err
	}
	smtps, err := models.SmtpCount()
	if err != nil {
		return nil, err
	}
	apiKeys, err := models.ApiKeyCount()
	if err != nil {
		return nil, err
	}
	return &types.DashboardDto{
		Templates: templates,
		Emails:    emails,
		Smtps:     smtps,
		ApiKeys:   apiKeys,
	}, nil
}
