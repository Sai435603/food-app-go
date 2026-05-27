package models

type OrderItem struct {
	BaseModel
	OrderId      string  `json:"order_id"`
	MenuId       string  `json:"menu_id"`
	Menu         Menu    `gorm:"foreignKey:MenuID" json:"menu,omitempty"`
	Quantity     int64   `json:"quantity"`
	PriceAtOrder float64 `json:"price_at_order"`
}
