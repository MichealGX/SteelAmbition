package battleRoom

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"platform/database"
)

func CreateRoom(c *gin.Context, db *gorm.DB, roomDTO RoomDTO, roomListDTO RoomListDTO) {
	// 插入新房间到 RoomList 表，生成房间ID
	roomList := database.RoomList{
		MaxNum:   roomListDTO.MaxNum,
		Num:      roomListDTO.Num,
		Status:   roomListDTO.Status,
		Survival: roomListDTO.Survival,
	}
	db = db.Session(&gorm.Session{NewDB: true})
	if err := db.Create(&roomList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to insert room"})
		return
	}

	// 创建新的数据表，例如 RoomData_123
	db = db.Session(&gorm.Session{NewDB: true})
	roomDataTableName := fmt.Sprintf("RoomData_%d", roomList.ID)
	if err := db.Table(roomDataTableName).AutoMigrate(&database.RoomData{}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to create room"})
		return
	}

	// 将 RoomID 存入 User 表
	db = db.Session(&gorm.Session{NewDB: true})
	if err := db.Table("users").Model(database.User{}).Where("id = ?", roomDTO.UserID).Update("room_id", roomList.ID).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to update roomID"})
		return
	}

	// 根据 vehicleID 获取 vehicleName
	var vehicleName string
	db = db.Session(&gorm.Session{NewDB: true})
	if err := db.Table("vehicles").Model(database.Vehicle{}).Where("id = ?", roomDTO.VehicleID).Pluck("vehicle_name", &vehicleName).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to get vehicleName"})
		return
	}
	roomDTO.VehicleName = vehicleName

	// 为新房间插入房主信息
	roomData := database.RoomData{
		UserID:      roomDTO.UserID,
		UserName:    roomDTO.UserName,
		VehicleID:   roomDTO.VehicleID,
		VehicleName: roomDTO.VehicleName,
		ReadyFlag:   roomDTO.ReadyFlag,
		Survive:     roomDTO.Survive,
	}
	err := InsertRoomRole(c, db, roomData, roomList.ID)
	if err != nil {
		return
	}

	roomArray := []database.RoomData{roomData}
	c.JSON(http.StatusCreated, gin.H{"code": 0, "msg": "success", "data": roomArray, "Num": roomListDTO.Num})
}
