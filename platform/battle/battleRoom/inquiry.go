package battleRoom

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"platform/database"
	"platform/users"
)

func Inquiry(c *gin.Context, db *gorm.DB, userDTO users.UserDTO, user *database.User, roomList *database.RoomList) error {
	// 从users获取用户信息
	db = db.Session(&gorm.Session{NewDB: true})
	if err := db.Model(database.User{}).Where("id = ?", userDTO.ID).First(user).Error; err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "User not found"})
		return err
	}
	if user.RoomID == 0 {
		c.JSON(400, gin.H{"code": 1, "msg": "User has not joined any room"})
		err := fmt.Errorf("User has not joined any room")
		return err
	}

	// 从roomList中获取房间信息
	db = db.Session(&gorm.Session{NewDB: true})
	if err := db.Model(database.RoomList{}).Where("id = ?", user.RoomID).First(roomList).Error; err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "Room not found"})
		return err
	}
	return nil
}
