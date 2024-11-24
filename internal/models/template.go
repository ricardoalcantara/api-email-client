package models

import "gorm.io/gorm"

type Template struct {
	gorm.Model
	Name         string `gorm:"size:255;not null;"`
	Slug         string `gorm:"size:255;not null;unique"`
	Subject      string `gorm:"type:text"`
	JsonSchema   string `gorm:"type:text;not null;"`
	TemplateHtml string `gorm:"type:text;not null;"`
	TemplateText string `gorm:"type:text;not null;"`
}

func (t *Template) Save() error {
	return db.Create(&t).Error
}

func (t *Template) Updates(update map[string]interface{}) error {
	return db.Model(&t).Updates(update).Error
}

func (t *Template) Update() error {
	return db.Model(&t).Save(&t).Error
}

func (t *Template) Delete() error {
	return db.Delete(&t).Error
}

func TemplateGet(id uint) (*Template, error) {
	var t Template
	err := db.Where("id = ?", id).Take(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func TemplateDelete(id uint) error {
	return db.Delete(&Template{}, id).Error
}

func TemplateList(pagination *Pagination) ([]Template, error) {
	var t []Template
	err := db.Scopes(pagination.GetScope).Find(&t).Error
	if err != nil {
		return nil, err
	}
	return t, nil
}

func TemplateGetBySlug(slug string) (*Template, error) {
	var t = Template{}
	err := db.Take(&t, "slug = ?", slug).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func TemplateCount() (int64, error) {
	var count int64
	err := db.Model(&Template{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
