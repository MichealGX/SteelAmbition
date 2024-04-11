package controllers

import (
	"fmt"
	"net/http"
	"platform/database"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"platform/battle/battleRoom"
	"platform/users"
	"platform/vehicle"
)

// RegisterHandler 处理注册用户请求
func RegisterHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userDTO users.UserDTO
		if err := c.ShouldBindJSON(&userDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Invalid request payload"})
			return
		}
		users.Register(c, db, userDTO)
	}
}

// LoginHandler 处理用户登录请求
func LoginHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userDTO users.UserDTO
		if err := c.ShouldBindJSON(&userDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Invalid request payload"})
			return
		}
		users.Login(c, db, userDTO)
	}
}

// AddVehicleHandler 处理添加战车请求
func AddVehicleHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var vehicleDTO vehicle.VehicleDTO
		if err := c.ShouldBindJSON(&vehicleDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Invalid request payload"})
			return
		}
		userDTO := users.ExtractUser(c)
		vehicleDTO.UserID = userDTO.ID
		vehicle.AddVehicle(c, db, vehicleDTO)
	}

}

// CreateRoomHandler 处理创建战斗房间请求
func CreateRoomHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userDTO := users.ExtractUser(c)

		var roomDTO battleRoom.RoomDTO
		var err error

		roomDTO.VehicleID, err = ExtractID(c, "vehicleID")
		if err != nil {
			return
		}

		roomDTO.ReadyFlag = false
		roomDTO.UserID = userDTO.ID
		roomDTO.UserName = userDTO.UserName

		var roomListDTO battleRoom.RoomListDTO
		roomListDTO.Num = 1
		if err := c.ShouldBindJSON(&roomListDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Invalid request payload"})
			return
		}
		battleRoom.CreateRoom(c, db, roomDTO, roomListDTO)
	}
}

// JoinRoomHandler 处理加入战斗房间请求
func JoinRoomHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userDTO := users.ExtractUser(c)
		var roomDTO battleRoom.RoomDTO
		roomDTO.UserID = userDTO.ID
		roomDTO.UserName = userDTO.UserName
		roomDTO.VehicleID, _ = ExtractID(c, "vehicleID")
		roomDTO.ReadyFlag = false

		roomid, _ := strconv.Atoi(c.Param("roomID"))
		roomID := uint(roomid)

		var vehicleName string
		if err := db.Model(&database.Vehicle{}).Select("vehicle_name").Where("id = ?", roomDTO.VehicleID).Scan(&vehicleName).Error; err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to get vehicleName"})
			return
		}
		roomDTO.VehicleName = vehicleName

		battleRoom.JoinRoom(c, db, roomDTO, roomID)
	}
}

// ReadyHandler 处理准备请求
func ReadyHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userDTO := users.ExtractUser(c)
		var roomID uint
		if err := db.Model(database.User{}).Where("id = ?", userDTO.ID).Select("room_id").Scan(&roomID).Error; err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to get roomID"})
			return
		}
		battleRoom.Ready(c, db, userDTO, roomID)
	}
}

// LeaveHandler 处理离开战斗房间请求
func LeaveHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userDTO := users.ExtractUser(c)
		battleRoom.Leave(c, db, userDTO)
	}
}
