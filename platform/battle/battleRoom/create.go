package battleRoom

import (
	"fmt"
	"net/http"
	"platform/database"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRoom(c *gin.Context, db *gorm.DB, roomDTO RoomDTO, roomListDTO RoomListDTO) {
	// 插入新房间到 RoomList 表，生成房间ID
	roomList := database.RoomList{
		MaxNum: roomListDTO.MaxNum,
		Num:    roomListDTO.Num,
		Status: 0,
	}
	if err := db.Create(&roomList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to insert room"})
		return
	}

	// 创建新的数据表，例如 RoomData_123
	roomDataTableName := fmt.Sprintf("RoomData_%d", roomList.ID)
	if err := db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY, created_at DATETIME, updated_at DATETIME, deleted_at DATETIME, user_id INT UNSIGNED, user_name VARCHAR(255), vehicle_id INT UNSIGNED, vehicle_name VARCHAR(255), ready_flag BOOLEAN)", roomDataTableName)).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to create room"})
		return
	}

	// 将 RoomID 存入 User 表
	if err := db.Model(&database.User{}).Where("id = ?", roomDTO.UserID).Update("room_id", roomList.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to update roomID"})
		return
	}

	// 根据 vehicleID 获取 vehicleName
	var vehicleName string
	if err := db.Model(&database.Vehicle{}).Where("id = ?", roomDTO.VehicleID).Pluck("vehicle_name", &vehicleName).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to get vehicleName"})
		return
	}
	roomDTO.VehicleName = vehicleName
	roomDTO.ReadyFlag = false

	// 为新房间插入房主信息
	roomData := database.RoomData{
		UserID:      roomDTO.UserID,
		UserName:    roomDTO.UserName,
		VehicleID:   roomDTO.VehicleID,
		VehicleName: roomDTO.VehicleName,
		ReadyFlag:   roomDTO.ReadyFlag,
	}
	roomData.CreatedAt = time.Now()
	roomData.UpdatedAt = time.Now()
	err := InsertRoomRole(c, db, roomData, roomList.ID)
	if err != nil {
		return
	}

	roomArray := []database.RoomData{roomData}
	c.JSON(http.StatusCreated, gin.H{"code": 0, "msg": "success", "data": roomArray, "Num": roomListDTO.Num})
}
