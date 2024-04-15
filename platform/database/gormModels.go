package database

import (
	"gorm.io/gorm"
	"time"
)

// User 用户表模型
type User struct {
	gorm.Model
	UserName string
	Password string
	Email    string
	RoomID   uint
}

// Vehicle 车辆表模型
type Vehicle struct {
	gorm.Model
	UserID                 uint
	VehicleName            string
	CoreModuleWeight       int
	CoreModuleEnergy       int
	AppearanceModuleWeight int
	WeaponModuleWeight     int
	WeaponModuleEnergy     int
	DefenseModuleWeight    int
	DefenseModuleEnergy    int
	WalkingModuleWeight    int
	WalkingModuleEnergy    int
	WalkingModuleSpeed     int
	Integral               int
	Number                 int
}

// BattleRecord 对战记录表模型
type BattleRecord struct {
	gorm.Model
	UserID    uint
	StartTime time.Time
	EndTime   time.Time
	Result    string
}

// RoomList 对战房间列表表模型
type RoomList struct {
	gorm.Model
	MaxNum      int
	Num         int
	Status      int
	Survival    int
	TimeLimit   int
	DamageValue int
}

// RoomData 对战房间数据表模型
type RoomData struct {
	gorm.Model
	UserID      uint
	UserName    string
	VehicleID   uint
	VehicleName string
	ReadyFlag   bool
	Survive     bool
}
