package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"platform/battle/battleRoom"
	"platform/users"
	"platform/vehicle"
	"strconv"
)

// RegisterHandler 处理注册用户请求
func RegisterHandler(Db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userDTO users.UserDTO
		if err := c.ShouldBindJSON(&userDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Invalid request payload"})
			return
		}
		users.Register(c, Db, userDTO)
	}
}

// LoginHandler 处理用户登录请求
func LoginHandler(Db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userDTO users.UserDTO
		if err := c.ShouldBindJSON(&userDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Invalid request payload"})
			return
		}
		users.Login(c, Db, userDTO)
	}
}

// AddVehicleHandler 处理添加战车请求
func AddVehicleHandler(Db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var vehicleDTO vehicle.VehicleDTO
		if err := c.ShouldBindJSON(&vehicleDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Invalid request payload"})
			return
		}
		userDTO := users.ExtractUser(c)
		vehicleDTO.UserID = userDTO.UserID
		vehicle.AddVehicle(c, Db, vehicleDTO)
	}

}

// CreateRoomHandler 处理创建战斗房间请求
func CreateRoomHandler(Db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var roomDTO battleRoom.RoomDTO
		roomDTO.VehicleID, _ = strconv.Atoi(c.Param("vehicleID"))
		var roomListDTO battleRoom.RoomListDTO
		if err := c.ShouldBindJSON(&roomListDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Invalid request payload"})
			return
		}
		battleRoom.Create(c, Db, roomDTO, roomListDTO)
	}
}
