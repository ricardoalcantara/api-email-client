package models

import "gorm.io/gorm"

type Template struct {
	gorm.Model
	Name         string `gorm:"size:255;not null;"`
	JsonSchema   string `gorm:"type:text"`
	Subject   string `gorm:"type:text"`
	TemplateHtml string `gorm:"type:text;not null;"`
	TemplateText string `gorm:"type:text;not null;"`
}

func (u *Template) Save() (*Template, error) {
	err := db.Create(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func TemplateGet(name string) (*Template, error) {
	var t = Template{Name: name}
	err := db.First(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}
