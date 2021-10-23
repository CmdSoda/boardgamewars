package bgw

type WeaponSystemCategory int

const (
	WeaponSystemCategoryNothing  WeaponSystemCategory = -1
	WeaponSystemCategoryGun      WeaponSystemCategory = 0
	WeaponSystemCategoryDropTank WeaponSystemCategory = 1
	WeaponSystemCategoryA2A      WeaponSystemCategory = 2
	WeaponSystemCategoryA2G      WeaponSystemCategory = 3
)

type WeaponSystem struct {
	WeaponSystemName string
	Category         string
	*Air2AirWeaponParameters
}

func GetWeaponSystemCategoryFromString(name string) WeaponSystemCategory {
	switch name {
	case "Gun":
		return WeaponSystemCategoryGun
	case "DropTank":
		return WeaponSystemCategoryDropTank
	case "A2A":
		return WeaponSystemCategoryA2A
	case "A2G":
		return WeaponSystemCategoryA2G
	}
	return WeaponSystemCategoryNothing
}

type WeaponSystemConfiguration struct {
	ConfigurationName string // Name of the configuration
	WeaponSystems     []WeaponSystem
}

type WeaponSystemConfigurationList []WeaponSystemConfiguration

func (wscl WeaponSystemConfigurationList) GetFromName(name string) *WeaponSystemConfiguration {
	for _, configuration := range wscl {
		if configuration.ConfigurationName == name {
			return &configuration
		}
	}
	return nil
}
