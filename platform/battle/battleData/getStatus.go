package battleData

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"platform/battle/battleRoom"
	"platform/database"
	"platform/users"
	"time"
)

func GetStatus(c *gin.Context, db *gorm.DB, rdb *redis.Client, userDTO users.UserDTO) {
	var user database.User
	var roomList database.RoomList
	err := battleRoom.Inquiry(c, db, userDTO, &user, &roomList)
	if err != nil {
		return
	}

	var statusDTO []StatusDTO
	ID, err := GetID(c, db, user)
	if err != nil {
		return
	}

	for i := 0; i < roomList.Num; i++ {
		// 获取用户的能量值数据
		userModule, err := GetEnergyData(c, rdb, ID[i].UserID)
		if err != nil {
			return
		}

		statusDTO = append(statusDTO, StatusDTO{
			UserID:              ID[i].UserID,
			VehicleID:           ID[i].VehicleID,
			CoreModuleEnergy:    userModule.CoreModuleEnergy,
			WeaponModuleEnergy:  userModule.WeaponModuleEnergy,
			DefenseModuleEnergy: userModule.DefenseModuleEnergy,
			WalkingModuleEnergy: userModule.WalkingModuleEnergy,
		})
	}

	timeRemaining, err, championID, championName := BattleFinishCheck(c, db, rdb, userDTO)
	if err != nil {
		return
	}
	timeRemaining = timeRemaining / time.Second
	// 从roomList中获取房间信息
	db = db.Session(&gorm.Session{NewDB: true})
	if err := db.Model(database.RoomList{}).Where("id = ?", user.RoomID).First(&roomList).Error; err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "Room not found"})
		return
	}

	if championID == 0 {
		c.JSON(200, gin.H{"code": 0, "msg": "Success", "data": statusDTO, "timeRemaining": timeRemaining, "battleStatus": roomList.Status})
	} else {
		c.JSON(200, gin.H{"code": 0, "msg": "Success", "data": statusDTO, "timeRemaining": timeRemaining, "battleStatus": roomList.Status, "championID": championID, "championName": championName})
	}

}
