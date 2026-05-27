package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeliveryPartner struct {
	Id               string  `gorm:"primaryKey;column:id" json:"id"`
	Name             string  `json:"name"`
	Rating           float64 `json:"rating"`
	VehicleNumber    string  `json:"vehicle_number"`
	CurrentLatitude  float64 `json:"current_latitude"`
	CurrentLongitude float64 `json:"current_longitude"`
}

func (deliveryPartner *DeliveryPartner) BeforeCreate(scope *gorm.DB) error {
	deliveryPartner.Id = uuid.New().String()
	return nil
}
