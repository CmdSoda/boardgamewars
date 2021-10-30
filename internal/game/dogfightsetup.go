package game

type DogfightSetup struct {
	TeamBlue AircraftIdList
	TeamRed  AircraftIdList
}

type DogfightGroup struct {
	BlueFighter *AircraftId
	BlueSupport *AircraftId // optional
	RedFighter  *AircraftId
	RedSupport  *AircraftId // optional
}

type Dogfight struct {
	Groups          []DogfightGroup
	TeamBlueWaiting AircraftIdList
	TeamRedWaiting  AircraftIdList
}
