package battleRoom

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"platform/database"
	"platform/users"
)

func SetDamage(c *gin.Context, db *gorm.DB, userDTO users.UserDTO, damageSet int) {
	var user database.User
	var roomList database.RoomList
	err := Inquiry(c, db, userDTO, &user, &roomList)
	if err != nil {
		return
	}

	// 更新房间对战伤害值
	db = db.Session(&gorm.Session{NewDB: true})
	if err := db.Model(database.RoomList{}).Where("id = ?", user.RoomID).Update("damage_value", damageSet).Error; err != nil {
		c.JSON(500, gin.H{"code": 1, "msg": "Failed to set damage"})
		return
	}

	c.JSON(200, gin.H{"code": 0, "msg": "Set damage successfully", "setterID": userDTO.ID, "setterName": userDTO.UserName, "damageSet": damageSet})
}
