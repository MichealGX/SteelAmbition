package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"platform/controllers"
)

// SetupRouter 设置路由
func SetupRouter(Db *sql.DB) *gin.Engine {
	// 创建一个默认的 Gin 引擎
	r := gin.Default()

	// 定义路由
	r.POST("/platform/users/register", controllers.RegisterHandler(Db))
	r.POST("/platform/users/login", controllers.LoginHandler(Db))

	// 返回路由引擎
	return r
}
