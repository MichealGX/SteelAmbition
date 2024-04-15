package vehicle

import (
	"net/http"
	"platform/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddVehicle(c *gin.Context, db *gorm.DB, vehicleDTO VehicleDTO) {
	// 创建车辆记录
	vehicle := database.Vehicle{
		UserID:                 vehicleDTO.UserID,
		VehicleName:            vehicleDTO.VehicleName,
		CoreModuleWeight:       vehicleDTO.CoreModuleWeight,
		CoreModuleEnergy:       vehicleDTO.CoreModuleEnergy,
		AppearanceModuleWeight: vehicleDTO.AppearanceModuleWeight,
		WeaponModuleWeight:     vehicleDTO.WeaponModuleWeight,
		WeaponModuleEnergy:     vehicleDTO.WeaponModuleEnergy,
		DefenseModuleWeight:    vehicleDTO.DefenseModuleWeight,
		DefenseModuleEnergy:    vehicleDTO.DefenseModuleEnergy,
		WalkingModuleWeight:    vehicleDTO.WalkingModuleWeight,
		WalkingModuleEnergy:    vehicleDTO.WalkingModuleEnergy,
		WalkingModuleSpeed:     vehicleDTO.WalkingModuleSpeed,
	}

	db = db.Session(&gorm.Session{NewDB: true})
	if err := db.Create(&vehicle).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to insert vehicle"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"code": 0, "msg": "success"})
}
