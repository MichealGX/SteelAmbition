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
	"time"
)

func Start(c *gin.Context, db *gorm.DB, rdb *redis.Client, userDTO users.UserDTO) {
	// 从users表中查询room_id
	var user database.User
	var roomList database.RoomList
	err := battleRoom.Inquiry(c, db, userDTO, &user, &roomList)
	if err != nil {
		return
	}

	// 检查房间是否满足开始游戏的条件
	err = battleRoom.BattleStartCheck(c, db, user.RoomID)
	if err != nil {
		return
	} else {
		// 更新房间状态
		db = db.Session(&gorm.Session{NewDB: true})
		if err := db.Model(database.RoomList{}).Where("id = ?", user.RoomID).Update("status", 1).Update("survival", roomList.Num).Error; err != nil {
			c.JSON(400, gin.H{"code": 1, "msg": "Failed to update room information"})
			return
		}

		// 生成战斗数据缓存,更新存活状态
		ID, err := GetID(c, db, user)
		if err != nil {
			return
		}

		ctx := context.Background()
		for i := 0; i < len(ID); i++ {
			var userModule database.UserModule
			//更新存活状态
			db = db.Session(&gorm.Session{NewDB: true})
			if err := db.Table(fmt.Sprintf("RoomData_%d", int(user.RoomID))).Where("user_id = ?", ID[i].UserID).Update("survive", true).Error; err != nil {
				c.JSON(400, gin.H{"code": 1, "msg": "Failed to update survival"})
				return
			}

			// 从vehicles中获取车辆信息
			db = db.Session(&gorm.Session{NewDB: true})
			if err = db.Table("vehicles").Model(database.UserModule{}).Where("id = ?", ID[i].VehicleID).First(&userModule).Error; err != nil {
				c.JSON(400, gin.H{"code": 1, "msg": "User modules not found"})
				return
			}

			// 将用户的能量值存储到 Redis 哈希中
			data, err := json.Marshal(userModule)
			if err != nil {
				c.JSON(400, gin.H{"code": 1, "msg": "Failed to marshal user modules"})
				return
			}
			err = rdb.HSet(ctx, fmt.Sprintf("user:%d", int(ID[i].UserID)), "modules", data).Err()
			if err != nil {
				c.JSON(400, gin.H{"code": 1, "msg": "Failed to set user modules"})
				return
			}
			// 设置过期时间
			err = rdb.Expire(ctx, fmt.Sprintf("user:%d", int(ID[i].UserID)), ConvertToDuration(2*roomList.TimeLimit)).Err()
			if err != nil {
				c.JSON(400, gin.H{"code": 1, "msg": "Failed to set expire time"})
				return
			}

		}

	}

	c.JSON(200, gin.H{"code": 0, "msg": "success", "roomID": user.RoomID})
}

// ConvertToDuration 将分钟数转换为时间间隔
func ConvertToDuration(n int) time.Duration {
	return time.Duration(n) * time.Minute
}
