package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"platform/controllers"
	"platform/users"
)

// SetupRouter 设置路由
func SetupRouter(Db *sql.DB) *gin.Engine {
	// 创建一个默认的 Gin 引擎
	r := gin.Default()

	// 定义路由
	r.POST("/platform/users/register", controllers.RegisterHandler(Db))
	r.POST("/platform/users/login", controllers.LoginHandler(Db))
	r.POST("/platform/vehicles/add", users.AuthMiddleware(), controllers.AddVehicleHandler(Db))
	r.POST("/platform/battle/:vehicleID/battleRoom/create", users.AuthMiddleware(), controllers.CreateRoomHandler(Db))

	// 返回路由引擎
	return r
}
