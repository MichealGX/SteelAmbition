package battleRoom

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"platform/database"
)

func UpdateNum(c *gin.Context, db *gorm.DB, roomID uint, num int) error {
	// 更新房间人数
	db = db.Session(&gorm.Session{NewDB: true})
	if err := db.Model(&database.RoomList{}).Where("id = ?", roomID).Update("num", num).Error; err != nil {
		fmt.Println("Failed to update roomList:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to update roomList"})
		return err
	}
	return nil
}
