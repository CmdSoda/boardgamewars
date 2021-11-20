package game

type Statistics struct {
	WeaponPerformance  WeaponPerformanceMap
	AircraftVsAircraft AircraftVersusPerformanceList
}

func NewStatistics() Statistics {
	s := Statistics{}
	s.WeaponPerformance = map[string]map[AircraftParametersId]map[string]*WeaponPerformanceStatistics{}
	s.AircraftVsAircraft = AircraftVersusPerformanceList{}
	return s
}
