package battleRoom

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"platform/database"
)

func BattleCheck(c *gin.Context, db *gorm.DB, roomID uint) (int, error) {
	var roomDataList []database.RoomData
	if err := db.Table(fmt.Sprintf("RoomData_%d", int(roomID))).Select("ready_flag").Find(&roomDataList).Error; err != nil {
		fmt.Println("Failed to query database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to query database"})
		return -1, err
	}

	for _, roomData := range roomDataList {
		if !roomData.ReadyFlag {
			return 0, nil
		}
	}

	var roomList database.RoomList
	if err := db.Model(database.RoomList{}).Where("id = ?", roomID).First(&roomList).Error; err != nil {
		fmt.Println("Failed to query database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to query database"})
		return -1, err
	}
	if roomList.Num < 2 {
		return 0, nil
	}

	return 1, nil
}
