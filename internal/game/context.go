package game

type Context struct {
	Configuration         Config
	AllAirbases           AirbasesMap
	AllAircraftParameters AircraftParametersMap
	Air2AirWeapons        Air2AirWeaponLibrary
	AllWarParties         WarPartyMap
	AllAircrafts          AircraftsMap
	AllPilots             PilotsMap
	AllCounters           CounterList
}

var Globals = Context{}
