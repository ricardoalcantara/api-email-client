package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:255;not null;"`
	Email    string `gorm:"size:255;not null;unique"`
	Password string `gorm:"size:255;not null;"`

	ApiKeys []ApiKey
}

func (u *User) Save() error {
	return db.Create(&u).Error
}

func (u *User) Update() error {
	return db.Save(&u).Error
}

func (u *User) Updates(update map[string]interface{}) error {
	return db.Model(&u).Updates(update).Error
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

var ErrMismatchedHashAndPassword = errors.New("invalid")

func (u *User) VerifyPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return ErrMismatchedHashAndPassword
		} else {
			return err
		}
	}

	return nil
}

func GetUserByEmail(email string) (*User, error) {
	u := User{Email: email}
	err := db.First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func GetUser(id uint) (*User, error) {
	u := User{Model: gorm.Model{ID: id}}
	err := db.First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}
