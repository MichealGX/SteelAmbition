package battleRoom

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"platform/database"
	"strconv"
)

func TraverseRoom(c *gin.Context, db *gorm.DB, roomID uint, roomArray *[]database.RoomData) error {
	// 遍历房间
	query := "RoomData_" + strconv.Itoa(int(roomID))
	err := db.Table(query).Model(database.RoomData{}).Find(roomArray).Error
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"code": 1, "msg": "Failed to traverse room"})
		return err
	}
	return nil
}
