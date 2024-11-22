package smtp

import (
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/pkg/types"
)

type SmtpService struct {
}

func NewSmtpService() *SmtpService {
	return &SmtpService{}
}

func (s *SmtpService) Create(smtpDto *types.CreateSmtpDto) (*types.SmtpDto, error) {
	smtp := models.Smtp{
		Name:     smtpDto.Name,
		Slug:     smtpDto.Slug,
		Server:   smtpDto.Server,
		Port:     smtpDto.Port,
		Email:    smtpDto.Email,
		User:     smtpDto.User,
		Password: smtpDto.Password,
		Default:  smtpDto.Default,
	}
	err := smtp.Save()
	if err != nil {
		return nil, err
	}

	view := NewSmtpDto(&smtp)
	return &view, nil
}

func (s *SmtpService) Get(slug string) (*types.SmtpDto, error) {
	smtp, err := models.SmtpGetBySlug(slug)
	if err != nil {
		return nil, err
	}

	view := NewSmtpDto(smtp)
	return &view, nil
}

func (s *SmtpService) List(pagination *models.Pagination) (*types.ListView[types.SmtpDto], error) {
	smtps, err := models.SmtpList(pagination)
	if err != nil {
		return nil, err
	}

	listView := types.ListView[types.SmtpDto]{}
	for _, t := range smtps {
		listView.List = append(listView.List, NewSmtpDto(&t))
	}

	return &listView, nil
}

func (s *SmtpService) Patch(slug string, updateSmtp *types.UpdateSmtpDto) (*types.SmtpDto, error) {
	smtp, err := models.SmtpGetBySlug(slug)
	if err != nil {
		return nil, err
	}

	update := map[string]interface{}{}
	if updateSmtp.Name != nil {
		update["name"] = *updateSmtp.Name
	}
	if updateSmtp.Slug != nil {
		update["slug"] = *updateSmtp.Slug
	}
	if updateSmtp.Server != nil {
		update["server"] = *updateSmtp.Server
	}
	if updateSmtp.Port != nil {
		update["port"] = *updateSmtp.Port
	}
	if updateSmtp.Email != nil {
		update["email"] = *updateSmtp.Email
	}
	if updateSmtp.User != nil {
		update["user"] = *updateSmtp.User
	}
	if updateSmtp.Password != nil && len(*updateSmtp.Password) > 0 {
		update["password"] = models.HashBase64Password(*updateSmtp.Password)
	}
	if updateSmtp.Default != nil {
		update["default"] = *updateSmtp.Default
	}

	if len(update) > 0 {
		err = smtp.Updates(update)
		if err != nil {
			return nil, err
		}
	}
	view := NewSmtpDto(smtp)
	return &view, nil
}

func (s *SmtpService) Update(slug string, updateSmtp *types.UpdateSmtpDto) (*types.SmtpDto, error) {
	smtp, err := models.SmtpGetBySlug(slug)
	if err != nil {
		return nil, err
	}

	smtp.Name = *updateSmtp.Name
	smtp.Slug = *updateSmtp.Slug
	smtp.Server = *updateSmtp.Server
	smtp.Port = *updateSmtp.Port
	smtp.Email = *updateSmtp.Email
	smtp.User = *updateSmtp.User
	smtp.Default = *updateSmtp.Default

	if updateSmtp.Password != nil && len(*updateSmtp.Password) > 0 {
		smtp.SetBase64Password(*updateSmtp.Password)
	}

	err = smtp.Update()
	if err != nil {
		return nil, err
	}

	view := NewSmtpDto(smtp)
	return &view, nil
}

func (s *SmtpService) Delete(slug string) error {
	smtp, err := models.SmtpGetBySlug(slug)
	if err != nil {
		return err
	}
	return smtp.Delete()
}
