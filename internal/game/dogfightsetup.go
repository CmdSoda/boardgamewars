package game

type DogfightSituation struct {
	BlueFighter Aircraft
	BlueSupport *Aircraft // optional
	RedFighter  Aircraft
	RedSupport  *Aircraft // optional
}

type DogfightSetup struct {
	TeamBlue   []Aircraft
	TeamRed    []Aircraft
	Situations []DogfightSituation
}

func NewDogfightSetup(blue []Aircraft, red []Aircraft) DogfightSetup {
	return DogfightSetup{
		TeamBlue:   blue,
		TeamRed:    red,
		Situations: []DogfightSituation{},
	}
}
