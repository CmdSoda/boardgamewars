package game

import "fmt"

// WeaponPerformanceMap ist eine Map auf Waffensystem-Name auf AircraftParametersId auf Config-Name
// Beispiel:
// m := WeaponPerformanceMap{}
// m["Aim-9"][acid]["Default"].Hit += 1
type WeaponPerformanceMap map[string]map[AircraftParametersId]map[string]*WeaponPerformanceStatistics

type WeaponPerformanceStatistics struct {
	Hit     int
	NotHit  int
	Damage1 int
	Damage2 int
	Damage3 int
}

func (wpm WeaponPerformanceMap) Dump() {
	fmt.Println("Statistics: Weapon Performance")
	for wname := range wpm {
		fmt.Println(wname)
		for acid := range wpm[wname] {
			acp := Globals.AllAircraftParameters[acid]
			fmt.Println("  " + acp.Name)
			for config := range wpm[wname][acid] {
				stats := wpm[wname][acid][config]
				count := stats.Damage1 + stats.Damage2 + stats.Damage3
				average := float32(stats.Damage1+stats.Damage2*2+stats.Damage3*3) /
					float32(count)
				sum := wpm[wname][acid][config].Hit + wpm[wname][acid][config].NotHit
				hitpc := float32(wpm[wname][acid][config].Hit) / float32(sum) * 100
				fmt.Printf("    %s Hit=%.1f%%, DmgAvg=%.1f (%d samples)\n", config, hitpc, average, sum)
			}
		}
	}
}

func (wpm *WeaponPerformanceMap) createMaps(weaponName string, acid AircraftId) {
	ac := Globals.AllAircrafts[acid]

	_, wnexist := (*wpm)[weaponName]
	if !wnexist {
		(*wpm)[weaponName] = map[AircraftParametersId]map[string]*WeaponPerformanceStatistics{}
	}

	_, apidexist := (*wpm)[weaponName][ac.AircraftParametersId]
	if !apidexist {
		(*wpm)[weaponName][ac.AircraftParametersId] = map[string]*WeaponPerformanceStatistics{}
	}

	_, cexist := (*wpm)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName]
	if !cexist {
		(*wpm)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName] = &WeaponPerformanceStatistics{}
	}
}

func (wpm *WeaponPerformanceMap) Hit(weaponName string, acid AircraftId) {
	wpm.createMaps(weaponName, acid)
	ac := Globals.AllAircrafts[acid]
	// Um 1 erhöhen.
	(*wpm)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].Hit =
		(*wpm)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].Hit + 1
}

func (wpm *WeaponPerformanceMap) NotHit(weaponName string, acid AircraftId) {
	wpm.createMaps(weaponName, acid)
	ac := Globals.AllAircrafts[acid]
	// Um 1 erhöhen.
	(*wpm)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].NotHit =
		(*wpm)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].NotHit + 1
}

func (wpm *WeaponPerformanceMap) Damage(weaponName string, acid AircraftId, dmg int) {
	wpm.createMaps(weaponName, acid)
	ac := Globals.AllAircrafts[acid]

	switch dmg {
	case 1:
		(*wpm)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].Damage1 =
			(*wpm)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].Damage1 + 1
	case 2:
		(*wpm)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].Damage2 =
			(*wpm)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].Damage2 + 1
	case 3:
		(*wpm)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].Damage3 =
			(*wpm)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].Damage3 + 1
	}
}
