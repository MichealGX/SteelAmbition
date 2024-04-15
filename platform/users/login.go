package users

import (
	"net/http"
	"platform/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context, db *gorm.DB, userDTO UserDTO) {
	// 查询用户
	var user database.User
	db = db.Session(&gorm.Session{NewDB: true})
	if err := db.Model(database.User{}).Where("user_name = ? AND password = ?", userDTO.UserName, userDTO.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 1, "msg": "Unauthorized"})
		return
	}

	// 生成 JWT
	token, err := SignedToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to sign token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "Login success", "token": token})
}
