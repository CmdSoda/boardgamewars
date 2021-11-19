package game

import "fmt"

// WeaponNameVsAircraftParameterIdMap ist eine Map auf Waffensystem-Name auf AircraftParametersId auf Config-Name
// Beispiel:
// m := WeaponNameVsAircraftParameterIdMap{}
// m["Aim-9"][acid]["Default"].Hit += 1
type WeaponNameVsAircraftParameterIdMap map[string]map[AircraftParametersId]map[string]*WeaponStatistics

type WeaponStatistics struct {
	Hit               int
	NotHit            int
}

type Statistics struct {
	W2A2C WeaponNameVsAircraftParameterIdMap
}

func NewStatistics() Statistics {
	s := Statistics{}
	s.W2A2C = map[string]map[AircraftParametersId]map[string]*WeaponStatistics{}
	return s
}

func (w2a2c WeaponNameVsAircraftParameterIdMap) Dump() {
	for wname := range w2a2c {
		fmt.Println(wname)
		for acid := range w2a2c[wname] {
			acp := Globals.AllAircraftParameters[acid]
			fmt.Println("  " + acp.Name)
			for config := range w2a2c[wname][acid] {
				sum := w2a2c[wname][acid][config].Hit + w2a2c[wname][acid][config].NotHit
				hitpc := float32(w2a2c[wname][acid][config].Hit) / float32(sum) * 100
				fmt.Printf("    %s Hit %.1f%% (%d samples)\n", config, hitpc, sum)
			}
		}
	}
}

func (w2a2c *WeaponNameVsAircraftParameterIdMap) createMaps(weaponName string, acid AircraftId) {
	ac := Globals.AllAircrafts[acid]

	_, wnexist := (*w2a2c)[weaponName]
	if !wnexist {
		(*w2a2c)[weaponName] = map[AircraftParametersId]map[string]*WeaponStatistics{}
	}

	_, apidexist := (*w2a2c)[weaponName][ac.AircraftParametersId]
	if !apidexist {
		(*w2a2c)[weaponName][ac.AircraftParametersId] = map[string]*WeaponStatistics{}
	}

	_, cexist := (*w2a2c)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName]
	if !cexist {
		(*w2a2c)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName] = &WeaponStatistics{}
	}
}

func (w2a2c *WeaponNameVsAircraftParameterIdMap) Hit(weaponName string, acid AircraftId) {
	w2a2c.createMaps(weaponName, acid)
	ac := Globals.AllAircrafts[acid]
	// Um 1 erhöhen.
	(*w2a2c)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].Hit =
		(*w2a2c)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].Hit +1
}

func (w2a2c *WeaponNameVsAircraftParameterIdMap) NotHit(weaponName string, acid AircraftId) {
	w2a2c.createMaps(weaponName, acid)
	ac := Globals.AllAircrafts[acid]
	// Um 1 erhöhen.
	(*w2a2c)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].NotHit =
		(*w2a2c)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].NotHit +1
}

