package game

// WeaponNameVsAircraftParameterIdMap ist eine Map auf Waffensystem-Name auf AircraftParametersId auf Config-Name
// Beispiel:
// m := WeaponNameVsAircraftParameterIdMap{}
// m["Aim-9"][acid]["Default"].Hit += 1
type WeaponNameVsAircraftParameterIdMap map[string]map[AircraftParametersId]map[string]WeaponStatistics

type WeaponStatistics struct {
	Hit               int
	NotHit            int
}

type Statistics struct {
	WeaponNameVsAircraftParameterIdMap
}
