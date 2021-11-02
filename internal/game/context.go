package game

type Context struct {
	AllAirbases           AirbasesMap
	AllAircraftParameters AircraftParametersMap
	Air2AirWeaponLibrary
	AllWarParties WarPartyMap
	AllAircrafts  AircraftsMap
	AllPilots     PilotsMap
}

var Globals = Context{}
