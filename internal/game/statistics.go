package game

type Statistics struct {
	WeaponPerformance  WeaponPerformanceMap
	AircraftVsAircraft AircraftVersusStatisticsList
}

func NewStatistics() Statistics {
	s := Statistics{}
	s.WeaponPerformance = map[string]map[AircraftParametersId]map[string]*WeaponPerformanceStatistics{}
	s.AircraftVsAircraft = AircraftVersusStatisticsList{}
	return s
}
