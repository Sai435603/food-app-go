package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	Id        string  `gorm:"primaryKey;column:id" json:"id"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func (address *Address) BeforeCreate(scope *gorm.DB) error {
	address.Id = uuid.New().String()
	return nil
}
