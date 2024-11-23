package auth

import (
	"github.com/google/uuid"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/internal/token"
	"github.com/ricardoalcantara/api-email-client/pkg/types"
)

type AuthService struct {
}

func (s *AuthService) Token(input types.TokenInput) (string, error) {

	client, err := models.GetUserByEmail(input.Email)
	if err != nil {
		return "", err
	}
	err = client.VerifyPassword(input.Password);
	if err != nil {
		return "", err
	}

	jti, _ := uuid.NewUUID()
	access_token, err := token.CreateAccessToken(client, jti)
	if err != nil {
		return "", err
	}

	return access_token, nil
}
