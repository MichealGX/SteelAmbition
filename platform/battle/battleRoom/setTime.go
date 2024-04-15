package battleRoom

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"platform/database"
	"platform/users"
)

func SetTime(c *gin.Context, db *gorm.DB, userDTO users.UserDTO, timeSet int) {
	// 从users表中查询room_id
	var user database.User
	var roomList database.RoomList
	db = db.Session(&gorm.Session{NewDB: true})
	err := Inquiry(c, db, userDTO, &user, &roomList)
	if err != nil {
		return
	}

	// 更新房间对战时长限制
	db = db.Session(&gorm.Session{NewDB: true})
	if err := db.Model(database.RoomList{}).Where("id = ?", user.RoomID).Update("time_limit", timeSet).Error; err != nil {
		c.JSON(500, gin.H{"code": 1, "msg": "Failed to set time"})
		return
	}

	c.JSON(200, gin.H{"code": 0, "msg": "Set time successfully", "setterID": userDTO.ID, "setterName": userDTO.UserName, "timeSet": timeSet})
}
