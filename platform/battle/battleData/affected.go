package battleData

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"platform/battle/battleRoom"
	"platform/database"
	"platform/users"
	"strconv"
)

func Affected(c *gin.Context, db *gorm.DB, rdb *redis.Client, userDTO users.UserDTO) {
	var user database.User
	var roomList database.RoomList
	err := battleRoom.Inquiry(c, db, userDTO, &user, &roomList)
	if err != nil {
		return
	}

	var roomData database.RoomData
	// 从roomData中获取战斗数据
	db = db.Session(&gorm.Session{NewDB: true})
	if err := db.Table(fmt.Sprintf("RoomData_%d", int(user.RoomID))).Where("user_id = ?", userDTO.ID).First(&roomData).Error; err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "Room data not found"})
		return
	}

	// 获取受击模块名称
	var affectedModule AffectedDTO
	if err := c.ShouldBindJSON(&affectedModule); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "Invalid request payload"})
		return
	}

	// 获取用户的能量值数据
	userModule, err := GetEnergyData(c, rdb, userDTO.ID)
	if err != nil {
		return
	}

	// 根据受击模块名称修改对应模块的能量值
	switch affectedModule.ModuleName {
	case "core":
		userModule.CoreModuleEnergy -= roomList.DamageValue
	case "weapon":
		userModule.WeaponModuleEnergy -= roomList.DamageValue
	case "defense":
		userModule.DefenseModuleEnergy -= roomList.DamageValue
	case "walking":
		userModule.WalkingModuleEnergy -= roomList.DamageValue
	default:
		c.JSON(400, gin.H{"code": 1, "msg": "Invalid module name"})
		return
	}

	// 确保能量值不会变成负数
	if userModule.CoreModuleEnergy < 0 {
		userModule.CoreModuleEnergy = 0
		//更新存活状态
		db = db.Session(&gorm.Session{NewDB: true})
		if err := db.Table(fmt.Sprintf("RoomData_%d", int(user.RoomID))).Where("user_id = ?", userDTO.ID).Update("survive", false).Error; err != nil {
			c.JSON(400, gin.H{"code": 1, "msg": "Failed to update survival"})
			return
		}
	}
	if userModule.WeaponModuleEnergy < 0 {
		userModule.WeaponModuleEnergy = 0
	}
	if userModule.DefenseModuleEnergy < 0 {
		userModule.DefenseModuleEnergy = 0
	}
	if userModule.WalkingModuleEnergy < 0 {
		userModule.WalkingModuleEnergy = 0
	}

	// 将修改后的能量值数据序列化为 JSON 格式
	updatedData, err := json.Marshal(userModule)
	if err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "Failed to marshal updated battle data"})
		return
	}

	// 将修改后的能量值数据存回 Redis 中
	ctx := context.Background()
	_, err = rdb.HSet(ctx, "user:"+strconv.Itoa(int(userDTO.ID)), "modules", updatedData).Result()
	if err != nil {
		c.JSON(500, gin.H{"code": 1, "msg": "Failed to set updated battle data in redis"})
		return
	}

	_, err, championID, championName := BattleFinishCheck(c, db, rdb, userDTO)
	if err != nil {
		return
	}
	// 从roomList中获取房间信息
	db = db.Session(&gorm.Session{NewDB: true})
	if err := db.Model(database.RoomList{}).Where("id = ?", user.RoomID).First(&roomList).Error; err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "Room not found"})
		return
	}

	// 返回战斗数据
	if championID == 0 {
		c.JSON(200, gin.H{"code": 0, "msg": "Success", "battleStatus": roomList.Status})
	} else {
		c.JSON(200, gin.H{"code": 0, "msg": "Success", "battleStatus": roomList.Status, "championID": championID, "championName": championName})
	}
}
