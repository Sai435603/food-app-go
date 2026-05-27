package models

type DeliveryPartner struct {
	BaseModel
	Name             string  `json:"name"`
	Rating           float64 `json:"rating"`
	VehicleNumber    string  `json:"vehicle_number"`
	CurrentLatitude  float64 `json:"current_latitude"`
	CurrentLongitude float64 `json:"current_longitude"`
}
