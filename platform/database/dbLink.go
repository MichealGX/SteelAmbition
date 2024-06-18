package database

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Link() (*gorm.DB, error) {
	// 连接MySQL服务器
	dsn := "user0:zx123456@tcp(121.36.4.215:3306)/SteelAmbition?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
