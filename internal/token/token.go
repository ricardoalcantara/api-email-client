package token

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/ricardoalcantara/api-email-client/internal/models"
)

type JwtCustomClaims struct {
	Name string
	jwt.RegisteredClaims
}

func CreateAccessToken(client *models.Client, jti uuid.UUID) (accessToken string, err error) {
	expiry, err := strconv.Atoi(os.Getenv("JWT_LIFESPAN"))
	if err != nil {
		return "", err
	}

	secret := os.Getenv("JWT_SECRET")
	exp := time.Now().Add(time.Minute * time.Duration(expiry))
	claims := &JwtCustomClaims{
		Name: client.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.Itoa(int(client.ID)),
			ExpiresAt: jwt.NewNumericDate(exp),
			ID:        jti.String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func CreateRefreshToken(client *models.Client, jti uuid.UUID) (refreshToken string, err error) {
	expiry, err := strconv.Atoi(os.Getenv("JWT_REFRESH_LIFESPAN"))
	if err != nil {
		return "", err
	}

	secret := os.Getenv("JWT_REFRES_TOKEN_SECRET")
	exp := time.Now().Add(time.Hour * time.Duration(expiry))
	claims := &JwtCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.Itoa(int(client.ID)),
			ExpiresAt: jwt.NewNumericDate(exp),
			ID:        jti.String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	token, err := ExtractToken(requestToken, secret)
	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, fmt.Errorf("invalid Token")
	}

	claims, err := ExtractClaims(token)

	if err != nil {
		return false, fmt.Errorf("invalid Token in claims")
	}

	if ok, err := IsBlacklisted(claims); ok {
		return false, fmt.Errorf("invalid Token")
	} else if err != nil {
		return false, fmt.Errorf("invalid Token in blacklist")
	}

	return true, nil
}

func ExtractToken(requestToken string, secret string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(requestToken, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func ExtractClaims(token *jwt.Token) (*JwtCustomClaims, error) {
	claims, ok := token.Claims.(*JwtCustomClaims)

	if !ok {
		return nil, fmt.Errorf("invalid Token")
	}
	return claims, nil
}

func IsBlacklisted(claims *JwtCustomClaims) (bool, error) {
	// key := "jti::black_list::" + claims.RegisteredClaims.ID

	// err := kv.Get(key).Err()
	// if err != nil && err != redis.Nil {
	// 	return true, err
	// } else if err == nil {
	// 	return true, nil
	// }

	return false, nil
}

func BlacklistIt(claims *JwtCustomClaims) {
	// key := "jti::black_list::" + claims.RegisteredClaims.ID
	// kv.Set(key, "0", time.Duration(time.Hour*24))
}
