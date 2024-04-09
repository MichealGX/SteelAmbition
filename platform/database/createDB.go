package database

import (
	"database/sql"
	"log"
)

func CreateDB(db *sql.DB) {
	// 创建数据库 "SteelAmbition"
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS SteelAmbition")
	if err != nil {
		log.Fatal(err)
	}

	// 使用 "SteelAmbition" 数据库
	_, err = db.Exec("USE SteelAmbition")
	if err != nil {
		log.Fatal(err)
	}

	// 创建用户表 "User"
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS User (
		UserID INT AUTO_INCREMENT PRIMARY KEY,
		Username VARCHAR(255) UNIQUE,
		Password VARCHAR(255),
		Email VARCHAR(255),
		VehicleID JSON
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// 创建车辆表 "Vehicle"
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Vehicle (
		VehicleID INT AUTO_INCREMENT PRIMARY KEY,
		UserID INT,
		VehicleName VARCHAR(255),
		CoreModule_Weight INT,
		CoreModule_Energy INT,
		AppearanceModule_Weight INT,
		WeaponModule_Weight INT,
		WeaponModule_Energy INT,
		DefenseModule_Weight INT,
		DefenseModule_Energy INT,
		WalkingModule_Weight INT,
		WalkingModule_Energy INT,
		WalkingModule_Speed INT,
		Integral INT,
		Number INT,
		FOREIGN KEY (UserID) REFERENCES User(UserID)
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// 创建对战记录表 "BattleRecord"
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS BattleRecord (
		RecordID INT AUTO_INCREMENT PRIMARY KEY,
		UserID INT,
		StartTime DATETIME,
		EndTime DATETIME,
		Result VARCHAR(255),
		FOREIGN KEY (UserID) REFERENCES User(UserID)
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// 创建对战房间列表表 "RoomList"
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS RoomList (
		RoomID INT AUTO_INCREMENT PRIMARY KEY,
		MaxNum INT
	)`)
	if err != nil {
		log.Fatal(err)
	}
}
