package token

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ricardoalcantara/api-email-client/internal/models"
)

type JwtCustomClaims struct {
	Name string
	jwt.RegisteredClaims
}

func CreateAccessToken(client *models.Client) (accessToken string, err error) {
	expiry, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	secret := os.Getenv("TOKEN_SECRET")

	if err != nil {
		return "", err
	}

	exp := time.Now().Add(time.Hour * time.Duration(expiry))
	claims := &JwtCustomClaims{
		Name: client.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			//  strconv.FormatUint(uint64(client.ID), 10)
			Subject:   base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(int(client.ID)))),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func IsAuthorized(requestToken string) (bool, error) {
	secret := os.Getenv("TOKEN_SECRET")
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractIDFromToken(requestToken string) (string, error) {
	secret := os.Getenv("TOKEN_SECRET")
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("invalid Token")
	}

	return claims["id"].(string), nil
}
