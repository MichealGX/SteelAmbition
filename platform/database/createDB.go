package database

import (
	"gorm.io/gorm"
)

// CreateDB 创建数据库表
func CreateDB(db *gorm.DB) {
	// 创建数据库 "SteelAmbition"
	if err := db.Exec("CREATE DATABASE IF NOT EXISTS SteelAmbition").Error; err != nil {
		panic(err)
	}

	// 使用数据库 "SteelAmbition"
	if err := db.Exec("USE SteelAmbition").Error; err != nil {
		panic(err)
	}

	// 迁移数据表
	err := db.AutoMigrate(&User{}, &Vehicle{}, &BattleRecord{}, &RoomList{})
	if err != nil {
		panic(err)
	}
}
