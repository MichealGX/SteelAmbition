package database

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Link() (*gorm.DB, error) {
	// 连接MySQL服务器
	dsn := "root:123456@tcp(localhost:3306)/?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
