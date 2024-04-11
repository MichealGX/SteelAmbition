package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"platform/controllers"
	"platform/users"
)

// SetupRouter 设置路由
func SetupRouter(Db *gorm.DB) *gin.Engine {
	// 创建一个默认的 Gin 引擎
	r := gin.Default()

	// 定义路由
	r.POST("/platform/users/register", controllers.RegisterHandler(Db))
	r.POST("/platform/users/login", controllers.LoginHandler(Db))
	r.POST("/platform/vehicles/add", users.AuthMiddleware(), controllers.AddVehicleHandler(Db))
	r.POST("/platform/battle/:vehicleID/battleRoom/create", users.AuthMiddleware(), controllers.CreateRoomHandler(Db))
	r.PUT("/platform/battle/:vehicleID/:roomID/battleRoom/join", users.AuthMiddleware(), controllers.JoinRoomHandler(Db))
	r.GET("/platform/battle/battleRoom/ready", users.AuthMiddleware(), controllers.ReadyHandler(Db))
	r.DELETE("/platform/battle/battleRoom/leave", users.AuthMiddleware(), controllers.LeaveHandler(Db))
	// 返回路由引擎
	return r
}
