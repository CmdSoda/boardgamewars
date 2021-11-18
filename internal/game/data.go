package game

type Data struct {
	Config
	AllAirbases           AirbasesMap
	AllAircraftParameters AircraftParametersMap
	Air2AirWeapons        Air2AirWeaponLibrary
	AllWarParties         WarPartyMap
	AllAircrafts          AircraftsMap
	AllPilots             PilotsMap
	AllCounters           CounterList
	Statistics
}

var Globals = Data{}
