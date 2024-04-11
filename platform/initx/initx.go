package initx

import (
	"gorm.io/gorm"
	"log"
	"platform/database"
	"platform/router"
)

var Db *gorm.DB

func Init() {
	// 链接数据库
	var err error
	Db, err = database.Link()
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}
	database.Cleanup(Db) //测试阶段清理数据库
	database.CreateDB(Db)
	// 设置路由
	r := router.SetupRouter(Db)

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Unable to start server:", err)
	}
}
