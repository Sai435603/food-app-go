package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Restaurant struct {
	Id     string `gorm:"primaryKey;column:id" json:"id"`
	Name   string `json:"name"`
	Rating string `json:"rating"`
}

func (restaurant *Restaurant) BeforeCreate(scope *gorm.DB) error {
	restaurant.Id = uuid.New().String()
	return nil
}
