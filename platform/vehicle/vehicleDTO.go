package vehicle

type VehicleDTO struct {
	UserID                  uint
	VehicleName             string `json:"vehicle_name" binding:"required"`
	CoreMoudle_Weight       int    `json:"core_w" binding:"required"`
	CoreMoudle_Energy       int    `json:"core_e" binding:"required"`
	AppearanceModule_Weight int    `json:"appearance_w"`
	WeaponModule_Weight     int    `json:"weapon_w" binding:"required"`
	WeaponModule_Energy     int    `json:"weapon_e" binding:"required"`
	DefenseModule_Weight    int    `json:"defense_w"`
	DefenseModule_Energy    int    `json:"defense_e"`
	WalkingModule_Weight    int    `json:"walking_w" binding:"required"`
	WalkingModule_Energy    int    `json:"walking_e" binding:"required"`
	WalkingModule_Speed     int    `json:"walking_s" binding:"required"`
	Integral                int
	Number                  int
}
