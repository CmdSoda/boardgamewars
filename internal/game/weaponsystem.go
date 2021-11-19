package game

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
	Depleted         bool // Weapon used, Tank has been droped.
	*Air2AirWeaponParameters
}

type WeaponSystemList []WeaponSystem

func CloneWeaponSystemList(acId AircraftParametersId, configname string) WeaponSystemList {
	for _, parameters := range Globals.AllAircraftParameters {
		if parameters.Id == acId {
			for _, configuration := range parameters.Configurations {
				if configuration.ConfigurationName == configname {
					wsl := make(WeaponSystemList, len(configuration.WeaponSystems))
					copy(wsl, configuration.WeaponSystems)
					return wsl
				}
			}
		}
	}
	return nil
}

func (ws *WeaponSystem) InitWeaponSystem() {
	switch GetWeaponSystemCategoryFromString(ws.Category) {
	case WeaponSystemCategoryA2A:
		aws, exist := GetAir2AirWeaponParametersFromName(ws.WeaponSystemName)
		if exist {
			ws.Air2AirWeaponParameters = &aws
		}
	}
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
	WeaponSystems     WeaponSystemList
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
