package models

import (
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

type Smtp struct {
	gorm.Model
	Name     string `gorm:"size:255;not null;"`
	Server   string `gorm:"size:255;not null;"`
	Port     uint16 `gorm:"not null;"`
	Email    string `gorm:"size:255;not null;"`
	User     string `gorm:"size:255;not null;"`
	Password string `gorm:"size:255;not null;"`
}

func (s *Smtp) Save() (*Smtp, error) {
	err := db.Create(&s).Error
	if err != nil {
		return &Smtp{}, err
	}
	return s, nil
}

func SmtpGet(name string) (*Smtp, error) {
	var t = Smtp{Name: name}
	err := db.First(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (s *Smtp) GetDialer() *gomail.Dialer {
	return gomail.NewDialer(s.Server, int(s.Port), s.User, s.Password)
}
