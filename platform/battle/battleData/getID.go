package battleData

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"platform/database"
)

func GetID(c *gin.Context, db *gorm.DB, user database.User) ([]IdDTO, error) {
	// 从RoomData表中获取用户ID和车辆ID
	db = db.Session(&gorm.Session{NewDB: true})
	var ID []IdDTO
	if err := db.Table(fmt.Sprintf("RoomData_%d", int(user.RoomID))).Select("user_id, vehicle_id").Find(&ID).Error; err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "Failed to get user and vehicle ID"})
		return nil, err
	}
	return ID, nil
}
