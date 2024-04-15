package users

import (
	"net/http"
	"platform/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(c *gin.Context, db *gorm.DB, userDTO UserDTO) {
	// 检查用户名是否已存在
	var count int64
	db = db.Session(&gorm.Session{NewDB: true})
	if err := db.Model(database.User{}).Where("user_name = ?", userDTO.UserName).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to query database"})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Username already exists"})
		return
	}

	// 插入新用户到数据库
	user := database.User{
		UserName: userDTO.UserName,
		Password: userDTO.Password,
		Email:    userDTO.Email,
	}
	db = db.Session(&gorm.Session{NewDB: true})
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to insert user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"code": 0, "msg": "success"})
}
