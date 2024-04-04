package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Link() (*sql.DB, error) {
	// 连接到 MySQL 数据库
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/SteelAmbition")
	if err != nil {
		return nil, err
	}
	return db, nil
}
