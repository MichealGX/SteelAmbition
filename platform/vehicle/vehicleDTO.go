package vehicle

type VehicleDTO struct {
	UserID                 uint
	VehicleName            string `json:"vehicle_name" binding:"required"`
	CoreModuleWeight       int    `json:"core_w" binding:"required"`
	CoreModuleEnergy       int    `json:"core_e" binding:"required"`
	AppearanceModuleWeight int    `json:"appearance_w"`
	WeaponModuleWeight     int    `json:"weapon_w" binding:"required"`
	WeaponModuleEnergy     int    `json:"weapon_e" binding:"required"`
	DefenseModuleWeight    int    `json:"defense_w"`
	DefenseModuleEnergy    int    `json:"defense_e"`
	WalkingModuleWeight    int    `json:"walking_w" binding:"required"`
	WalkingModuleEnergy    int    `json:"walking_e" binding:"required"`
	WalkingModuleSpeed     int    `json:"walking_s" binding:"required"`
	Integral               int
	Number                 int
}
