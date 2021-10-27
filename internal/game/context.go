package game

type Context struct {
	PilotRoster
	AirbaseList
	AircraftLibrary
	Air2AirWeaponLibrary
}

var Globals = Context{}
