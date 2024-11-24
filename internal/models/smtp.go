package models

import (
	"encoding/base64"
	"errors"

	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

type Smtp struct {
	gorm.Model
	Name     string `gorm:"size:255;not null;"`
	Slug     string `gorm:"size:255;not null;unique"`
	Server   string `gorm:"size:255;not null;"`
	Port     uint16 `gorm:"smallint;not null;"`
	Email    string `gorm:"size:255;not null;"`
	User     string `gorm:"size:255;not null;"`
	Password string `gorm:"size:255;not null;"`
	Default  bool   `gorm:"boolean;not null;default:0"`
}

func (s *Smtp) Save() error {
	return db.Create(&s).Error
}

func (t *Smtp) Updates(update map[string]interface{}) error {
	return db.Model(&t).Updates(update).Error
}

func (t *Smtp) Update() error {
	return db.Model(&t).Save(&t).Error
}

func (t *Smtp) Delete() error {
	return db.Delete(&t).Error
}

func (s *Smtp) SetBase64Password(password string) {
	s.Password = base64.StdEncoding.EncodeToString([]byte(password))
}

func HashBase64Password(password string) string {
	return base64.StdEncoding.EncodeToString([]byte(password))
}

func (s *Smtp) GetBase64Password() (string, error) {
	res, err := base64.StdEncoding.DecodeString(s.Password)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func (s *Smtp) Base64Password() {
	s.Password = base64.StdEncoding.EncodeToString([]byte(s.Password))
}

func SmtpList(pagination *Pagination) ([]Smtp, error) {
	var s []Smtp
	err := db.Scopes(pagination.GetScope).Find(&s).Error
	if err != nil {
		return nil, err
	}
	return s, nil
}

func SmtpGet(id uint) (*Smtp, error) {
	var s = Smtp{}
	err := db.Take(&s, id).Error
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func SmtpGetByName(name string) (*Smtp, error) {
	var s = Smtp{}
	err := db.Take(&s, "name = ?", name).Error
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func SmtpGetFirst() (*Smtp, error) {
	var s = Smtp{}
	err := db.First(&s).Error
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func SmtpGetDefault() (*Smtp, error) {
	var s = Smtp{}
	err := db.Take(&s, "`default` = ?", true).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return SmtpGetFirst()
		}

		return nil, err
	}
	return &s, nil
}

func SmtpDisableDefault() error {
	return db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&Smtp{}).Update("default", false).Error
}

func SmtpDeleteById(id uint) error {
	return db.Delete(&Smtp{}, id).Error
}

func SmtpGetBySlug(slug string) (*Smtp, error) {
	var t = Smtp{}
	err := db.Take(&t, "slug = ?", slug).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (s *Smtp) GetDialer() (*gomail.Dialer, error) {
	password, err := s.GetBase64Password()
	if err != nil {
		return nil, err
	}
	return gomail.NewDialer(s.Server, int(s.Port), s.User, password), nil
}

func SmtpCount() (int64, error) {
	var count int64
	err := db.Model(&Smtp{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
