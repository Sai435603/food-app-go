package models

// Orders table
type Order struct {
	BaseModel
	UserId            string          `json:"user_id"`
	User              User            `gorm:"foreignKey:UserID"`
	RestaurantId      string          `json:"restaurant_id"`
	Restaurant        Restaurant      `gorm:"foreignKey:RestaurantID"`
	DeliveryPartnerId string          `json:"delivery_partner_id"`
	DeliveryPartner   DeliveryPartner `gorm:"foreignKey:DeliveryPartnerID" json:"delivery_partner,omitempty"`
	AddressId         string          `json:"address_id"`
	Address           Address         `gorm:"foreignKey:AddressID" json:"address,omitempty"`
	OrderStatus       string          `json:"order_status"`
	OrderItems        []OrderItem     `gorm:"foreignKey:OrderID"`
}
