package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"platform/users"
)

// RegisterHandler 处理注册用户请求
func RegisterHandler(Db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userDTO users.UserDTO
		if err := c.ShouldBindJSON(&userDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Invalid request payload"})
			return
		}
		users.Register(c, Db, userDTO)
	}
}

// LoginHandler 处理用户登录请求
func LoginHandler(Db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userDTO users.UserDTO
		if err := c.ShouldBindJSON(&userDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Invalid request payload"})
			return
		}
		users.Login(c, Db, userDTO)
	}
}
