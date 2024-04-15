package router

import (
	"github.com/gin-gonic/gin"
	"platform/controllers"
	"platform/users"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	// 创建一个默认的 Gin 引擎
	r := gin.Default()

	// 定义路由
	r.POST("/platform/users/register", controllers.RegisterHandler())
	r.POST("/platform/users/login", controllers.LoginHandler())

	r.POST("/platform/vehicles/add", users.AuthMiddleware(), controllers.AddVehicleHandler())

	r.POST("/platform/battle/:vehicleID/room/create", users.AuthMiddleware(), controllers.CreateRoomHandler())
	r.PUT("/platform/battle/:vehicleID/:roomID/room/join", users.AuthMiddleware(), controllers.JoinRoomHandler())
	r.GET("/platform/battle/room/ready", users.AuthMiddleware(), controllers.ReadyHandler())
	r.DELETE("/platform/battle/room/leave", users.AuthMiddleware(), controllers.LeaveHandler())
	r.POST("/platform/battle/room/:time_limit/set_time", users.AuthMiddleware(), controllers.SetTimeHandler())
	r.POST("/platform/battle/room/set_damage/:damage_value", users.AuthMiddleware(), controllers.SetDamageHandler())

	r.GET("/platform/battle/data/start", users.AuthMiddleware(), controllers.StartHandler())
	r.POST("platform/battle/data/affected", users.AuthMiddleware(), controllers.AffectedHandler())
	r.GET("/platform/vehicles/data/status", users.AuthMiddleware(), controllers.StatusHandler())

	// 返回路由引擎
	return r
}
