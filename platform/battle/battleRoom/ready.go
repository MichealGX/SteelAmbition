package battleRoom

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"platform/database"
	"platform/users"
)

func Ready(c *gin.Context, db *gorm.DB, userDTO users.UserDTO, roomID uint) {
	var roomData database.RoomData
	db = db.Session(&gorm.Session{NewDB: true})
	if err := db.Table(fmt.Sprintf("RoomData_%d", int(roomID))).Model(database.RoomData{}).Where("user_id = ?", userDTO.ID).First(&roomData).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to get roomData"})
		return
	}

	// 切换准备状态
	roomData.ReadyFlag = !roomData.ReadyFlag
	// 更新用户准备状态
	db = db.Session(&gorm.Session{NewDB: true})
	if err := db.Table(fmt.Sprintf("RoomData_%d", int(roomID))).Save(&roomData).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to update readyFlag"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "ReadyFlag": roomData.ReadyFlag})
}
