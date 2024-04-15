package controllers

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"net/http"
	"platform/battle/battleData"
	"platform/database"
	"strconv"

	"github.com/gin-gonic/gin"
	"platform/battle/battleRoom"
	"platform/users"
	"platform/vehicle"
)

// RegisterHandler 处理注册用户请求
func RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := database.SteelAmbitionLink()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to connect to database"})
			return
		}

		var userDTO users.UserDTO
		if err := c.ShouldBindJSON(&userDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Invalid request payload"})
			return
		}
		users.Register(c, db, userDTO)
	}
}

// LoginHandler 处理用户登录请求
func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := database.SteelAmbitionLink()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to connect to database"})
			return
		}

		var userDTO users.UserDTO
		if err := c.ShouldBindJSON(&userDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Invalid request payload"})
			return
		}
		users.Login(c, db, userDTO)
	}
}

// AddVehicleHandler 处理添加战车请求
func AddVehicleHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := database.SteelAmbitionLink()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to connect to database"})
			return
		}

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
func CreateRoomHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := database.SteelAmbitionLink()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to connect to database"})
			return
		}

		userDTO := users.ExtractUser(c)
		var roomDTO battleRoom.RoomDTO

		roomDTO.VehicleID, err = ExtractID(c, "vehicleID")
		if err != nil {
			return
		}

		roomDTO.ReadyFlag = false
		roomDTO.Survive = true
		roomDTO.UserID = userDTO.ID
		roomDTO.UserName = userDTO.UserName

		var roomListDTO battleRoom.RoomListDTO
		roomListDTO.Num = 1
		roomListDTO.Survival = 1
		roomListDTO.Status = 0
		if err := c.ShouldBindJSON(&roomListDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Invalid request payload"})
			return
		}
		battleRoom.CreateRoom(c, db, roomDTO, roomListDTO)
	}
}

// JoinRoomHandler 处理加入战斗房间请求
func JoinRoomHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := database.SteelAmbitionLink()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to connect to database"})
			return
		}

		userDTO := users.ExtractUser(c)
		var roomDTO battleRoom.RoomDTO
		roomDTO.UserID = userDTO.ID
		roomDTO.UserName = userDTO.UserName
		roomDTO.VehicleID, err = ExtractID(c, "vehicleID")
		if err != nil {
			return
		}
		roomDTO.ReadyFlag = false
		roomDTO.Survive = true

		roomId, _ := strconv.Atoi(c.Param("roomID"))
		roomID := uint(roomId)

		var vehicleName string
		if err := db.Table("vehicles").Model(database.Vehicle{}).Select("vehicle_name").Where("id = ?", roomDTO.VehicleID).Scan(&vehicleName).Error; err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to get vehicleName"})
			return
		}
		roomDTO.VehicleName = vehicleName

		battleRoom.JoinRoom(c, db, roomDTO, roomID)
	}
}

// ReadyHandler 处理准备请求
func ReadyHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := database.SteelAmbitionLink()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to connect to database"})
			return
		}

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
func LeaveHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := database.SteelAmbitionLink()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to connect to database"})
			return
		}

		userDTO := users.ExtractUser(c)
		battleRoom.Leave(c, db, userDTO)
	}
}

// SetTimeHandler 处理设置战斗时间请求
func SetTimeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := database.SteelAmbitionLink()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to connect to database"})
			return
		}

		userDTO := users.ExtractUser(c)
		timeSet, err := strconv.Atoi(c.Param("time_limit"))
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Invalid time"})
			return
		}
		// 检查时间是否合法
		if timeSet <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Invalid time"})
			return
		}

		battleRoom.SetTime(c, db, userDTO, timeSet)
	}
}

// SetDamageHandler 处理设置战斗伤害请求
func SetDamageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := database.SteelAmbitionLink()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to connect to database"})
			return
		}

		userDTO := users.ExtractUser(c)
		damageSet, err := strconv.Atoi(c.Param("damage_value"))
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Invalid damage"})
			return
		}
		// 检查伤害是否合法
		if damageSet <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Invalid damage"})
			return
		}

		battleRoom.SetDamage(c, db, userDTO, damageSet)
	}
}

// StartHandler 处理开始战斗请求
func StartHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := database.SteelAmbitionLink()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to connect to database"})
			return
		}
		var rdb *redis.Client
		database.RedisLink(&rdb)

		userDTO := users.ExtractUser(c)
		battleData.Start(c, db, rdb, userDTO)
	}
}

// AffectedHandler 处理车辆受击数据请求
func AffectedHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := database.SteelAmbitionLink()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to connect to database"})
			return
		}
		var rdb *redis.Client
		database.RedisLink(&rdb)

		userDTO := users.ExtractUser(c)
		battleData.Affected(c, db, rdb, userDTO)
	}
}

// StatusHandler 处理获取车辆状态数据请求
func StatusHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := database.SteelAmbitionLink()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to connect to database"})
			return
		}
		var rdb *redis.Client
		database.RedisLink(&rdb)

		userDTO := users.ExtractUser(c)
		battleData.GetStatus(c, db, rdb, userDTO)
	}
}
