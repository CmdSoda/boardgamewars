package game

type DogfightSetup struct {
	TeamBlue []AircraftId
	TeamRed  []AircraftId
}

type DogfightGroup struct {
	BlueFighter AircraftId
	BlueSupport *AircraftId // optional
	RedFighter  AircraftId
	RedSupport  *AircraftId // optional
}

type Dogfight struct {
	Groups          []DogfightGroup
	TeamBlueWaiting []AircraftId
	TeamRedWaiting  []AircraftId
}
