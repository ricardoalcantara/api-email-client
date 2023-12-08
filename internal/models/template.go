package models

import "gorm.io/gorm"

type Template struct {
	gorm.Model
	Name         string `gorm:"size:255;not null;"`
	Subject      string `gorm:"type:text"`
	TemplateHtml string `gorm:"type:text;not null;"`
	TemplateText string `gorm:"type:text;not null;"`
}

func (u *Template) Save() error {
	return db.Create(&u).Error
}

func TemplateGet(id uint) (*Template, error) {
	var t Template
	err := db.Where("id = ?", id).Take(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func TemplateList(pagination *Pagination) ([]Template, error) {
	var t []Template
	err := db.Scopes(pagination.GetScope).Find(&t).Error
	if err != nil {
		return nil, err
	}
	return t, nil
}

func TemplateGetById(name string) (*Template, error) {
	var t = Template{Name: name}
	err := db.Take(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}
