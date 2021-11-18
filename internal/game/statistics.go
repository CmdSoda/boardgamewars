package game

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
	WeaponNameVsAircraftParameterIdMap
}

func NewStatistics() Statistics {
	s := Statistics{}
	s.WeaponNameVsAircraftParameterIdMap = map[string]map[AircraftParametersId]map[string]*WeaponStatistics{}
	return s
}

func (s *Statistics) Hit(weaponName string, acid AircraftId) {
	ac := Globals.AllAircrafts[acid]

	_, wnexist := s.WeaponNameVsAircraftParameterIdMap[weaponName]
	if !wnexist {
		s.WeaponNameVsAircraftParameterIdMap[weaponName] = map[AircraftParametersId]map[string]*WeaponStatistics{}
	}

	_, apidexist := s.WeaponNameVsAircraftParameterIdMap[weaponName][ac.AircraftParametersId]
	if !apidexist {
		s.WeaponNameVsAircraftParameterIdMap[weaponName][ac.AircraftParametersId] = map[string]*WeaponStatistics{}
	}

	_, cexist := s.WeaponNameVsAircraftParameterIdMap[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName]
	if !cexist {
		s.WeaponNameVsAircraftParameterIdMap[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName] = &WeaponStatistics{}
	}

	// Um 1 erhöhen.
	s.WeaponNameVsAircraftParameterIdMap[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].Hit =
		s.WeaponNameVsAircraftParameterIdMap[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].Hit +1
}

func (s *Statistics) NotHit(weaponName string, acid AircraftId) {
	ac := Globals.AllAircrafts[acid]

	_, wnexist := s.WeaponNameVsAircraftParameterIdMap[weaponName]
	if !wnexist {
		s.WeaponNameVsAircraftParameterIdMap[weaponName] = map[AircraftParametersId]map[string]*WeaponStatistics{}
	}

	_, apidexist := s.WeaponNameVsAircraftParameterIdMap[weaponName][ac.AircraftParametersId]
	if !apidexist {
		s.WeaponNameVsAircraftParameterIdMap[weaponName][ac.AircraftParametersId] = map[string]*WeaponStatistics{}
	}

	_, cexist := s.WeaponNameVsAircraftParameterIdMap[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName]
	if !cexist {
		s.WeaponNameVsAircraftParameterIdMap[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName] = &WeaponStatistics{}
	}

	// Um 1 erhöhen.
	s.WeaponNameVsAircraftParameterIdMap[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].NotHit =
		s.WeaponNameVsAircraftParameterIdMap[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].NotHit +1
}

