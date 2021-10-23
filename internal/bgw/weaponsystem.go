package bgw

type WeaponSystemCategory uint

const (
	DropTank WeaponSystemCategory = 0
	A2A                           = 1
	Gun                           = 2
	A2G                           = 3
)

type WeaponSystem struct {
	EquipmentId
	Category WeaponSystemCategory
}

type WeaponSystemConfiguration struct {
	ConfigurationName string // Name of the configuration
	WeaponSystems     []WeaponSystem
}
