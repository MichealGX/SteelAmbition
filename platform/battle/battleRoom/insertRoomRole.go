package battleRoom

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"platform/database"
)

func InsertRoomRole(c *gin.Context, db *gorm.DB, roomData database.RoomData, roomID uint) error {
	// 插入战斗房间角色
	//query := fmt.Sprintf("INSERT INTO RoomData_%d (user_id, user_name, vehicle_id, vehicle_name, ready_flag) VALUES (?, ?, ?, ?, ?)", roomData.ID)
	//err := db.Exec(query, roomData.UserID, roomData.UserName, roomData.VehicleID, roomData.VehicleName, roomData.ReadyFlag).Error
	query := fmt.Sprintf("RoomData_%d", int(roomID))
	err := db.Table(query).Create(&roomData).Error
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to insert roomData"})
		return err
	}
	return nil
}
