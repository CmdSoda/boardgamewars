package game

type Context struct {
	AirbaseList
	AircraftLibrary
	Air2AirWeaponLibrary
	AllWarParties WarPartyList
	AllAircrafts map[AircraftId]Aircraft
	AllPilots map[PilotId]Pilot
}

var Globals = Context{}
