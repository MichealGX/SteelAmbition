package battleData

type AffectedDTO struct {
	ModuleName string `json:"moduleName"`
}

type StatusDTO struct {
	UserID              uint `json:"userID"`
	VehicleID           uint `json:"vehicleID"`
	CoreModuleEnergy    int  `json:"coreModuleEnergy"`
	WeaponModuleEnergy  int  `json:"weaponModuleEnergy"`
	DefenseModuleEnergy int  `json:"defenseModuleEnergy"`
	WalkingModuleEnergy int  `json:"walkingModuleEnergy"`
}

type IdDTO struct {
	UserID    uint
	VehicleID uint
}
