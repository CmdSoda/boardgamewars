package game

type Context struct {
	AirbaseList
	AircraftLibrary
	Air2AirWeaponLibrary
	WarPartyList
}

var Globals = Context{}
