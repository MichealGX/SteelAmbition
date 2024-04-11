package battleRoom

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"platform/database"
	"platform/users"
)

func Leave(c *gin.Context, db *gorm.DB, userDTO users.UserDTO) {
	// 从users表中查询room_id
	var user database.User
	if err := db.Model(database.User{}).Where("id = ?", userDTO.ID).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "User not found"})
		return
	}
	if user.RoomID == 0 {
		c.JSON(400, gin.H{"code": 1, "msg": "User has not joined any room"})
		return
	}

	// 从房间中删除用户
	var roomData database.RoomData
	result := db.Table(fmt.Sprintf("RoomData_%d", int(user.RoomID))).Model(database.RoomData{}).Where("user_id = ?", userDTO.ID).First(&roomData)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(400, gin.H{"code": 1, "msg": "Failed to leave room"})
		return
	}
	err := db.Table(fmt.Sprintf("RoomData_%d", int(user.RoomID))).Model(database.RoomData{}).Delete(&roomData).Error
	if err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "Failed to leave room"})
		return
	}

	// 查询房间人数
	var roomList database.RoomList
	if err := db.Model(database.RoomList{}).Where("id = ?", user.RoomID).First(&roomList).Error; err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "Failed to retrieve room information"})
		return
	}

	// 更新房间人数
	if err := db.Model(database.RoomList{}).Where("id = ?", user.RoomID).Update("num", roomList.Num-1).Error; err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "Failed to update room information"})
		return
	}

	// 房间对战状态检查
	battleStatus, err := BattleCheck(c, db, user.RoomID)
	if err != nil {
		return
	}

	// 如果房间人数为0，删除房间
	if roomList.Num-1 == 0 {
		var roomlistD database.RoomList
		if err := db.Model(database.RoomList{}).Where("id = ?", user.RoomID).First(&roomlistD).Error; err != nil {
			c.JSON(400, gin.H{"code": 1, "msg": "Failed to retrieve room information"})
			return
		}
		if err := db.Model(database.RoomList{}).Where("id = ?", user.RoomID).Delete(&roomlistD).Error; err != nil {
			c.JSON(400, gin.H{"code": 1, "msg": "Failed to delete room"})
			return
		}
	}

	// 更新users表中的room_id
	if err := db.Model(database.User{}).Where("id = ?", userDTO.ID).Update("room_id", 0).Error; err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "Failed to update user information"})
		return
	}

	// 成功离开房间
	c.JSON(200, gin.H{"code": 0, "msg": "Successfully left the room", "num": roomList.Num - 1, "BattleStatus": battleStatus})
}
