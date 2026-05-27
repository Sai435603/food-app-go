package models

type Address struct {
	BaseModel
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
