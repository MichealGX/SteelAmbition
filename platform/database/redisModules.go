package database

type UserModule struct {
	CoreModuleEnergy    int `gorm:"column:core_module_energy"`
	WeaponModuleEnergy  int `gorm:"column:weapon_module_energy"`
	DefenseModuleEnergy int `gorm:"column:defense_module_energy"`
	WalkingModuleEnergy int `gorm:"column:walking_module_energy"`
}
