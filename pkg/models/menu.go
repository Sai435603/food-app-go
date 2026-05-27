package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Menu struct {
	Id           string     `gorm:"primaryKey;column:id" json:"id"`
	RestaurantId string     `json:"restaurant_id"`
	Restaurant   Restaurant `gorm:"foreignKey:RestaurantID" json:"-"`
	ItemName     string     `json:"item_name"`
	Price        float64    `json:"price"`
	Category     string     `json:"category"`
	IsAvailable  bool       `json:"is_available"`
}

func (menu *Menu) BeforeCreate(scope *gorm.DB) error {
	menu.Id = uuid.New().String()
	return nil
}
