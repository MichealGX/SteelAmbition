package battleRoom

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"platform/database"
)

func CheckNum(c *gin.Context, db *gorm.DB, roomID uint) int {
	var roomList database.RoomList
	if err := db.Select("num").Where("id = ?", roomID).First(&roomList).Error; err != nil {
		fmt.Println("Failed to get roomID:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to get roomID"})
		return -1
	}
	return roomList.Num
}
