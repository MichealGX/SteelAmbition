package battleRoom

import "gorm.io/gorm"

type RoomDTO struct {
	gorm.Model
	UserID      uint
	UserName    string `json:"UserName"`
	VehicleID   uint   `json:"VehicleID"`
	VehicleName string `json:"VehicleName"`
	ReadyFlag   bool   `json:"ReadyFlag"`
}

type RoomListDTO struct {
	gorm.Model
	MaxNum int `json:"maxNum"`
	Num    int
}
