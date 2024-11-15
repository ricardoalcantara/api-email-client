package models

import (
	"time"

	"gorm.io/gorm"
)

type ApiKey struct {
	gorm.Model
	Name        string `gorm:"size:255;not null;"`
	KeyHash     string `gorm:"size:255;not null;unique"`
	LastUsed    *time.Time
	IpWhitelist string `gorm:"size:255;not null;"`
	ExpiresAt   *time.Time
	UserId      uint

	User *User
}

func (a *ApiKey) Save() error {
	return db.Create(&a).Error
}

func (a *ApiKey) Delete() error {
	return db.Delete(&a).Error
}

func GetApiKeyByHash(keyHash string) (*ApiKey, error) {
	a := ApiKey{KeyHash: keyHash}
	err := db.First(&a).Error
	if err != nil {
		return nil, err
	}
	return &a, nil
}
