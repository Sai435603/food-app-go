package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id       string `gorm:"primaryKey;column:id" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func (user *User) BeforeCreate(scope *gorm.DB) error {
	user.Id = uuid.New().String()
	return nil
}
