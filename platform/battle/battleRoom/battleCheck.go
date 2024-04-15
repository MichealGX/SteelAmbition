package battleRoom

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"platform/database"
)

func BattleStartCheck(c *gin.Context, db *gorm.DB, roomID uint) error {
	// 查询房间中的所有用户的准备状态
	var roomDataList []database.RoomData
	db = db.Session(&gorm.Session{NewDB: true})
	if err := db.Table(fmt.Sprintf("RoomData_%d", int(roomID))).Select("ready_flag").Find(&roomDataList).Error; err != nil {
		fmt.Println("Failed to query database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to query database"})
		return err
	}

	for _, roomData := range roomDataList {
		if !roomData.ReadyFlag {
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Some users are not ready"})
			err := fmt.Errorf("some users are not ready")
			return err
		}
	}

	// 查询房间中的用户数量
	var roomList database.RoomList
	db = db.Session(&gorm.Session{NewDB: true})
	if err := db.Model(database.RoomList{}).Where("id = ?", roomID).First(&roomList).Error; err != nil {
		fmt.Println("Failed to query database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to query database"})
		return err
	}
	if roomList.Num < 2 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Insufficient number of people"})
		err := fmt.Errorf("insufficient number of people")
		return err
	}

	// 检查是否设置了战斗时间
	if roomList.TimeLimit == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Battle time not set"})
		err := fmt.Errorf("battle time not set")
		return err
	}
	// 检查是否设置了战斗伤害
	if roomList.DamageValue == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Battle damage not set"})
		err := fmt.Errorf("battle damage not set")
		return err
	}

	return nil
}
