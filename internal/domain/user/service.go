package user

import (
	"errors"

	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/pkg/types"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) UpdatePassword(userId uint, input types.UpdatePasswordDto) error {
	user, err := models.GetUser(userId)
	if err != nil {
		return err
	}

	err = user.VerifyPassword(input.CurrentPassword)
	if err != nil {
		return errors.New(err.Error() + " password")
	}

	user.SetPassword(input.NewPassword)
	if err = user.Update(); err != nil {
		return err
	}
	return nil
}
