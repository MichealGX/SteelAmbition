package database

import (
	"gorm.io/gorm"
	"log"
)

func Cleanup(db *gorm.DB) {
	// 使用 "SteelAmbition" 数据库
	err := db.Exec("USE SteelAmbition").Error
	if err != nil {
		log.Fatal(err)
	}
	// 清理数据库
	err = db.Migrator().DropTable(&Vehicle{}, &BattleRecord{}, &User{}, &RoomList{}, &RoomData{})
	if err != nil {
		panic(err)
	}
	db.Exec("DROP TABLE IF EXISTS RoomData_1")
}
