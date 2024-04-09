package battleRoom

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context, Db *sql.DB, roomDTO RoomDTO, roomListDTO RoomListDTO) {
	// 插入新房间到RoomList表,生成房间ID
	result, err := Db.Exec("INSERT INTO RoomList (MaxNum) VALUES (?)", roomListDTO.MaxNum)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to insert room"})
		return
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to get roomID"})
		return
	}
	// 将获取到的 ID 赋值给 roomListDTO.RoomID
	roomListDTO.RoomID = int(lastInsertID)

	// 创建新的数据表，例如 RoomData_123
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS RoomData_%d (UserID INT, UserName VARCHAR(255), VehicleID INT, VehicleName VARCHAR(255), ReadyFlag BOOLEAN, PRIMARY KEY (UserID))", roomListDTO.RoomID)
	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to create room"})
		return
	}
	//根据vehicleID获取vehicleName
	err = Db.QueryRow("SELECT VehicleName FROM Vehicle WHERE VehicleID = ?", roomDTO.VehicleID).Scan(&roomDTO.VehicleName)
	//为新房间插入房主信息
	_, err = Db.Exec(fmt.Sprintf("INSERT INTO RoomData_%d (UserID, UserName, VehicleID, VehicleName, ReadyFlag) VALUES (?, ?, ?, ?, ?)", roomListDTO.RoomID), roomDTO.UserID, roomDTO.UserName, roomDTO.VehicleID, roomDTO.VehicleName, false)
	c.JSON(http.StatusCreated, gin.H{"code": 0, "msg": "success"})
}
