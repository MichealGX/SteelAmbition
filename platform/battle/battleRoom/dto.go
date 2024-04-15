package battleRoom

import (
	"gorm.io/gorm"
)

type RoomDTO struct {
	gorm.Model
	UserID      uint
	UserName    string `json:"UserName"`
	VehicleID   uint   `json:"VehicleID"`
	VehicleName string `json:"VehicleName"`
	ReadyFlag   bool   `json:"ReadyFlag"`
	Survive     bool   `json:"Survive"`
}

type RoomListDTO struct {
	gorm.Model
	MaxNum      int `json:"maxNum"`
	Num         int
	Status      int
	Survival    int
	TimeLimit   int
	DamageValue int
}
