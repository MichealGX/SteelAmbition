package initx

import (
	"gorm.io/gorm"
	"log"
	"platform/database"
	"platform/router"
)

// Init 初始化
func Init() {
	var db *gorm.DB

	// 链接mysql数据库,初始化数据库
	var err error
	db, err = database.Link()
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}
	//database.CleanUp(db) //测试阶段清理数据库
	database.CreateDB(db)

	// 设置路由
	r := router.SetupRouter()

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Unable to start server:", err)
	}
}
