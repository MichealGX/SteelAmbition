package battleRoom

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"platform/database"
	"strconv"
)

func JoinRoom(c *gin.Context, db *gorm.DB, roomDTO RoomDTO, roomID uint) {
	// 在进行下一次查询前，清除字段选择
	db = db.Session(&gorm.Session{NewDB: true})
	// 或者使用 Select("*") 来重置字段选择
	// db = db.Select("*")

	// 检查房间是否存在
	var roomList database.RoomList
	if err := db.Table("room_lists").Model(database.RoomList{}).First(&roomList, roomID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Room does not exist"})
		return
	}

	// 检查userID是否已经在房间中
	db = db.Session(&gorm.Session{NewDB: true})
	var count int64
	if err := db.Table("RoomData_"+strconv.Itoa(int(roomList.ID))).Where("user_id = ?", roomDTO.UserID).Count(&count).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Failed to query database"})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "User already in room"})
		return
	}

	// 查询房间最大人数限制
	maxNum := roomList.MaxNum

	// 检查房间人数
	num := CheckNum(c, db, roomID)
	if num == -1 {
		return
	} else if num == maxNum {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Room is full"})
		return
	}
	num++

	// 更新房间人数
	if err := UpdateNum(c, db, roomID, num); err != nil {
		return
	}

	// 加入房间
	roomData := database.RoomData{
		UserID:      roomDTO.UserID,
		UserName:    roomDTO.UserName,
		VehicleID:   roomDTO.VehicleID,
		VehicleName: roomDTO.VehicleName,
		ReadyFlag:   roomDTO.ReadyFlag,
		Survive:     roomDTO.Survive,
	}
	err := InsertRoomRole(c, db, roomData, roomID)
	if err != nil {
		return
	}

	// 将 RoomID 存入 User 表
	if err := db.Model(&database.User{}).Where("id = ?", roomDTO.UserID).Update("room_id", roomID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to update roomID"})
		return
	}

	// 返回房间信息
	var roomArray []database.RoomData
	err = TraverseRoom(c, db, roomID, &roomArray)
	if err != nil {
		return
	}

	c.JSON(http.StatusCreated, gin.H{"code": 0, "msg": "success", "data": roomArray, "Num": num})
}
