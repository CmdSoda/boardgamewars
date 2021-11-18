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

func (wmap WeaponNameVsAircraftParameterIdMap) Dump() {
	for wname := range wmap {
		fmt.Println(wname)
		for acid := range wmap[wname] {
			acp := Globals.AllAircraftParameters[acid]
			fmt.Println("  " + acp.Name)
			for config := range wmap[wname][acid] {
				hitpc := float32(wmap[wname][acid][config].Hit) /
					(float32(wmap[wname][acid][config].Hit) + float32(wmap[wname][acid][config].NotHit)) * 100
				fmt.Printf("    %s Hit%% %f\n", config, hitpc)
			}
		}
	}
}

func (s *Statistics) Hit(weaponName string, acid AircraftId) {
	ac := Globals.AllAircrafts[acid]

	_, wnexist := s.W2A2C[weaponName]
	if !wnexist {
		s.W2A2C[weaponName] = map[AircraftParametersId]map[string]*WeaponStatistics{}
	}

	_, apidexist := s.W2A2C[weaponName][ac.AircraftParametersId]
	if !apidexist {
		s.W2A2C[weaponName][ac.AircraftParametersId] = map[string]*WeaponStatistics{}
	}

	_, cexist := s.W2A2C[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName]
	if !cexist {
		s.W2A2C[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName] = &WeaponStatistics{}
	}

	// Um 1 erhöhen.
	s.W2A2C[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].Hit =
		s.W2A2C[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].Hit +1
}

func (s *Statistics) NotHit(weaponName string, acid AircraftId) {
	ac := Globals.AllAircrafts[acid]

	_, wnexist := s.W2A2C[weaponName]
	if !wnexist {
		s.W2A2C[weaponName] = map[AircraftParametersId]map[string]*WeaponStatistics{}
	}

	_, apidexist := s.W2A2C[weaponName][ac.AircraftParametersId]
	if !apidexist {
		s.W2A2C[weaponName][ac.AircraftParametersId] = map[string]*WeaponStatistics{}
	}

	_, cexist := s.W2A2C[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName]
	if !cexist {
		s.W2A2C[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName] = &WeaponStatistics{}
	}

	// Um 1 erhöhen.
	s.W2A2C[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].NotHit =
		s.W2A2C[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].NotHit +1
}

