package models

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Name         string `gorm:"size:255;not null;"`
	ClientId     string `gorm:"size:255;not null;unique"`
	ClientSecret string `gorm:"size:255;not null;"`
}

func (u *Client) Save() (*Client, error) {
	err := db.Create(&u).Error
	if err != nil {
		return &Client{}, err
	}
	return u, nil
}

func LoginCheck(clientId string, clientSecret string) (*Client, error) {

	var err error

	c := Client{}

	err = db.Where("client_id = ? AND client_secret = ?", clientId, clientSecret).Take(&c).Error

	if err != nil {
		return nil, err
	}

	return &c, nil
}
