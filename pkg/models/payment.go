package models

type Payment struct {
	BaseModel
	OrderId        string  `json:"order_id"`
	Order          Order   `gorm:"foreignKey:OrderID" json:"-"`
	Total          float64 `json:"total"`
	PaymentsMethod string  `json:"payment_method"`
	Status         string  `json:"status"`
}
