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
		CoreModuleWeight:       vehicleDTO.CoreMoudle_Weight,
		CoreModuleEnergy:       vehicleDTO.CoreMoudle_Energy,
		AppearanceModuleWeight: vehicleDTO.AppearanceModule_Weight,
		WeaponModuleWeight:     vehicleDTO.WeaponModule_Weight,
		WeaponModuleEnergy:     vehicleDTO.WeaponModule_Energy,
		DefenseModuleWeight:    vehicleDTO.DefenseModule_Weight,
		DefenseModuleEnergy:    vehicleDTO.DefenseModule_Energy,
		WalkingModuleWeight:    vehicleDTO.WalkingModule_Weight,
		WalkingModuleEnergy:    vehicleDTO.WalkingModule_Energy,
		WalkingModuleSpeed:     vehicleDTO.WalkingModule_Speed,
	}

	if err := db.Create(&vehicle).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to insert vehicle"})
		return
	}

	//// 构建更新的 JSON 数组
	//updatedJSON := fmt.Sprintf("json_array_append(vehicle_id, '$', '%d')", vehicle.ID)

	//// 更新用户的车辆信息
	//if err := db.Model(database.User{}).Where("id = ?", vehicleDTO.UserID).Update("vehicle_id", gorm.Expr(updatedJSON)).Error; err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to update user's vehicle info"})
	//	return
	//}

	c.JSON(http.StatusCreated, gin.H{"code": 0, "msg": "success"})
}
