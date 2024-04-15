package battleRoom

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"platform/database"
	"platform/users"
)

func Leave(c *gin.Context, db *gorm.DB, userDTO users.UserDTO) {
	var user database.User
	var roomList database.RoomList
	err := Inquiry(c, db, userDTO, &user, &roomList)
	if err != nil {
		return
	}

	// 从房间中删除用户
	db = db.Session(&gorm.Session{NewDB: true})
	err = db.Table(fmt.Sprintf("RoomData_%d", int(user.RoomID))).Unscoped().Where("user_id = ?", userDTO.ID).Delete(nil).Error
	if err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "Failed to leave room"})
		return
	}

	// 更新房间人数
	db = db.Session(&gorm.Session{NewDB: true})
	if err := db.Model(database.RoomList{}).Where("id = ?", user.RoomID).Update("num", roomList.Num-1).Error; err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "Failed to update room information"})
		return
	}

	// 如果房间人数为0，删除房间
	if roomList.Num-1 == 0 {
		db = db.Session(&gorm.Session{NewDB: true})
		if err := db.Model(database.RoomList{}).Unscoped().Where("id = ?", user.RoomID).Delete(nil).Error; err != nil {
			c.JSON(400, gin.H{"code": 1, "msg": "Failed to delete room"})
			return
		}
		// 创建表名
		tableName := fmt.Sprintf("RoomData_%d", user.RoomID)
		// 构建 DROP TABLE 语句
		dropTableSQL := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", tableName)
		// 执行原生 SQL 语句删除表
		err := db.Exec(dropTableSQL).Error
		if err != nil {
			c.JSON(500, gin.H{"code": 1, "msg": "Failed to drop table"})
			return
		}
	}

	// 更新users表中的room_id
	db = db.Session(&gorm.Session{NewDB: true})
	if err := db.Model(database.User{}).Where("id = ?", userDTO.ID).Update("room_id", 0).Error; err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "Failed to update user information"})
		return
	}

	// 返回房间信息
	var roomArray []database.RoomData
	if roomList.Num-1 > 0 {
		err = TraverseRoom(c, db, user.RoomID, &roomArray)
		if err != nil {
			return
		}
	}

	// 成功离开房间
	c.JSON(200, gin.H{"code": 0, "msg": "Successfully left the room", "data": roomArray, "num": roomList.Num - 1})
}
