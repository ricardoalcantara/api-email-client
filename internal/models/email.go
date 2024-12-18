package models

import (
	"time"

	"gorm.io/gorm"
)

type Email struct {
	gorm.Model
	SentAt   *time.Time
	SmtpId   uint
	To       string `gorm:"size:255;not null;"`
	Subject  string `gorm:"size:255;not null;"`
	HtmlBody string `gorm:"type:text;null;"`
	TextBody string `gorm:"type:text;null;"`

	Smtp *Smtp
}

func (e *Email) Save() error {
	return db.Create(&e).Error
}

func EmailUpdateSent(id uint) error {
	return db.Model(&Email{}).Where("id = ?", id).Update("sent_at", time.Now()).Error
}

func EmailGet(id uint) (*Email, error) {
	e := Email{Model: gorm.Model{ID: id}}
	err := db.Joins("Smtp").Find(&e).Error
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func EmailList(pagination *Pagination) ([]Email, error) {
	var s []Email
	err := db.Scopes(pagination.GetScope).
		Unscoped().
		Joins("Smtp").
		Omit("HtmlBody", "TextBody").
		Find(&s, "emails.deleted_at IS NULL").Error
	if err != nil {
		return nil, err
	}
	return s, nil
}

func EmailCount() (int64, error) {
	var count int64
	err := db.Model(&Email{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
