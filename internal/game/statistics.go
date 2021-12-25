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

func FreeForAll() {
	for _, parameters := range Globals.AllAircraftParameters {
		// Blau erstellen
		acblue := NewAircraft(parameters.Name, "Default", WarPartyIdUSA)
		acblue.FillSeatsWithNewPilots(OF1)
		for _, parameters2 := range Globals.AllAircraftParameters {
			if parameters.Name != parameters2.Name {
				// Rot erstellen
				acred := NewAircraft(parameters2.Name, "Default", WarPartyIdRussia)
				acred.FillSeatsWithNewPilots(OF1)
				for i := 0; i < 10000; i++ {
					acblue.ReviveAndRepair()
					acblue.Rearm()
					ds := NewDogfightSetup()
					ds.AddBlue(acblue.AircraftId)
					// Blau reparieren und bewaffnen
					ds.AddRed(acred.AircraftId)
					// Simulation
					d := NewDogfight(ds)
					d.DistributeAircraftsToGroups()
					d.Simulate()
				}
			}
		}
	}
}
