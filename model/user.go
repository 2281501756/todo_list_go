package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `gorm:"serializer:json"`
	UserName   string `gorm:"unique" json:"user_name"`
	Password   string `json:"password"`
}

func (this *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	this.Password = string(bytes)
	return nil
}

func (this *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(this.Password), []byte(password))
	return err == nil
}
