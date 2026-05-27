package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	Id             string  `gorm:"primaryKey;column:id" json:"id"`
	OrderId        string  `json:"order_id"`
	Order          Order   `gorm:"foreignKey:OrderID" json:"-"`
	Total          float64 `json:"total"`
	PaymentsMethod string  `json:"payment_method"`
	Status         string  `json:"status"`
}

func (payment *Payment) BeforeCreate(scope *gorm.DB) error {
	payment.Id = uuid.New().String()
	return nil
}
