package models

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ApiKey struct {
	gorm.Model
	Name        string `gorm:"size:255;not null;"`
	Prefix      string `gorm:"size:255;not null;"`
	Key         string `gorm:"size:255;not null;"`
	LastUsed    *time.Time
	IpWhitelist string `gorm:"size:255;not null;"`
	ExpiresAt   *time.Time
	UserId      uint

	User *User
}

func GenerateApiKey() (*ApiKey, string, error) {
	prefix, key, hashedKey, err := generateAPIKey()
	if err != nil {
		return nil, "", err
	}

	apiKey := ApiKey{
		Prefix: prefix,
		Key:    hashedKey,
	}

	fullKey := prefix + "." + key

	return &apiKey, fullKey, nil
}

func (a *ApiKey) Save() error {
	return db.Create(&a).Error
}

func (a *ApiKey) Delete() error {
	return db.Delete(&a).Error
}

func ApiKeyGet(userId, id uint) (*ApiKey, error) {
	var s = ApiKey{}
	err := db.Take(&s, "user_id = ? AND id = ?", userId, id).Error
	if err != nil {
		return nil, err
	}
	return &s, nil
}

var ErrMismatchedHashAndKey = errors.New("invalid api key")

func ApiKeyGetByHash(hash string) (*ApiKey, error) {
	parts := strings.Split(hash, ".")
	if len(parts) != 2 {
		return nil, errors.New("invalid hash")
	}
	var apiKey = ApiKey{}
	err := db.Take(&apiKey, "prefix = ?", parts[0]).Error
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(apiKey.Key), []byte(parts[1]))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, ErrMismatchedHashAndKey
		} else {
			return nil, err
		}
	}

	return &apiKey, nil
}

func ApiKeyList(userId uint, pagination *Pagination) ([]ApiKey, error) {
	var s []ApiKey
	err := db.Scopes(pagination.GetScope).Find(&s, "user_id = ?", userId).Error
	if err != nil {
		return nil, err
	}
	return s, nil
}

func ApiKeyDelete(userId, id uint) error {
	return db.Delete(&ApiKey{}, "user_id = ? AND id = ?", userId, id).Error
}

func generateAPIKey() (string, string, string, error) {
	prefix := make([]byte, 12)
	if _, err := rand.Read(prefix); err != nil {
		return "", "", "", err
	}
	prefixStr := base64.URLEncoding.EncodeToString(prefix)[:12]

	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return "", "", "", err
	}
	keyStr := base64.URLEncoding.EncodeToString(key)[:32]

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(keyStr), bcrypt.DefaultCost)
	if err != nil {
		return "", "", "", err
	}

	return prefixStr, keyStr, string(hashedBytes), nil
}

func ApiKeyCount() (int64, error) {
	var count int64
	err := db.Model(&ApiKey{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
