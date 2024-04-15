package battleData

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"platform/battle/battleRoom"
	"platform/database"
	"platform/users"
	"time"
)

func BattleFinishCheck(c *gin.Context, db *gorm.DB, rdb *redis.Client, userDTO users.UserDTO) (time.Duration, error, uint, string) {
	// 检查战斗是否结束
	var user database.User
	var roomList database.RoomList
	var championID uint
	var championName string
	err := battleRoom.Inquiry(c, db, userDTO, &user, &roomList)
	if err != nil {
		return -1, err, 0, ""
	}

	//检查存活人数
	var count int = 0
	ID, err := GetID(c, db, user)
	if err != nil {
		return -1, err, 0, ""
	}

	for i := 0; i < len(ID); i++ {
		var roomData database.RoomData
		db = db.Session(&gorm.Session{NewDB: true})
		if err := db.Table(fmt.Sprintf("RoomData_%d", int(user.RoomID))).Where("user_id = ?", ID[i].UserID).First(&roomData).Error; err != nil {
			c.JSON(400, gin.H{"code": 1, "msg": "Room data not found"})
			return -1, err, 0, ""
		}
		if roomData.Survive == true {
			count++
			if count == 1 {
				championID = ID[i].UserID
				db = db.Session(&gorm.Session{NewDB: true})
				if err = db.Table("users").Model(database.User{}).Select("user_name").Where("id = ?", ID[i].UserID).First(&championName).Error; err != nil {
					c.JSON(400, gin.H{"code": 1, "msg": "User not found"})
					return -1, err, 0, ""
				}
			}
		}
	}

	//检查剩余时间
	key := fmt.Sprintf("user:%d", int(userDTO.ID))
	ttl, err := GetTTL(c, rdb, key)
	if err != nil {
		return -1, err, 0, ""
	}

	if count == 1 || ttl <= ConvertToDuration(roomList.TimeLimit) {
		//战斗结束
		//更新房间状态
		db = db.Session(&gorm.Session{NewDB: true})
		if err := db.Model(database.RoomList{}).Where("id = ?", user.RoomID).Update("status", 0).Error; err != nil {
			c.JSON(400, gin.H{"code": 1, "msg": "Failed to update room information"})
			return 0, err, championID, championName
		}

		//积分规则
		//...

		return 0, nil, championID, championName
	}

	return ttl - ConvertToDuration(roomList.TimeLimit), nil, 0, ""
}
