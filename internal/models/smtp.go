package models

import "gorm.io/gorm"

type Smtp struct {
	gorm.Model
	Name     string `gorm:"size:255;not null;"`
	Server   string `gorm:"size:255;not null;"`
	Port     uint16 `gorm:"not null;"`
	Tls      bool   `gorm:"not null;"`
	User     string `gorm:"size:255;not null;"`
	Password string `gorm:"size:255;not null;"`
}

func (u *Smtp) SaveSmtp() (*Smtp, error) {
	err := db.Create(&u).Error
	if err != nil {
		return &Smtp{}, err
	}
	return u, nil
}
