package vehicle

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddVehicle(c *gin.Context, Db *sql.DB, vehicleDTO VehicleDTO) {
	// 插入车辆信息
	_, err := Db.Exec(`INSERT INTO Vehicle (
                     UserID,
                     VehicleName,
                     CoreModule_Weight,
                     CoreModule_Energy,
                     AppearanceModule_Weight,
                     WeaponModule_Weight,
                     WeaponModule_Energy,
                     DefenseModule_Weight,
                     DefenseModule_Energy,
                     WalkingModule_Weight,
                     WalkingModule_Energy,
                     WalkingModule_Speed
                     ) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`,
		vehicleDTO.UserID,
		vehicleDTO.VehicleName,
		vehicleDTO.CoreMoudle_Weight,
		vehicleDTO.CoreMoudle_Energy,
		vehicleDTO.AppearanceModule_Weight,
		vehicleDTO.WeaponModule_Weight,
		vehicleDTO.WeaponModule_Energy,
		vehicleDTO.DefenseModule_Weight,
		vehicleDTO.DefenseModule_Energy,
		vehicleDTO.WalkingModule_Weight,
		vehicleDTO.WalkingModule_Energy,
		vehicleDTO.WalkingModule_Speed)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to insert vehicle"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"code": 0, "msg": "success"})
}
