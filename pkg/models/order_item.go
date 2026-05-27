package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderItem struct {
	Id           string  `gorm:"primaryKey;column:id" json:"id"`
	OrderId      string  `json:"order_id"`
	MenuId       string  `json:"menu_id"`
	Menu         Menu    `gorm:"foreignKey:MenuID" json:"menu,omitempty"`
	Quantity     int64   `json:"quantity"`
	PriceAtOrder float64 `json:"price_at_order"`
}

func (order_item *OrderItem) BeforeCreate(scope *gorm.DB) error {
	order_item.Id = uuid.New().String()
	return nil
}
