package game

type World struct {
	CurrentStep       StepTime
	Airbases          []AirbaseId
	AircraftsInTheAir []AircraftId
}

func NewWorld() World {
	return World{Airbases: []AirbaseId{}}
}

func (w *World) AddAirbase(st AirbaseId) {
	//var s Stepper = st
	w.Airbases = append(w.Airbases, st)
}

func (w *World) MoveAircraftToAir(abid AirbaseId, acid AircraftId) {
}

func (w *World) MoveAircraftInAirToAirbase(acid AircraftId, abid AirbaseId) {
}

func (w *World) Step(st StepTime) {
	w.CurrentStep = w.CurrentStep + st

	// Step all Airbases
	for _, abid := range w.Airbases {
		ab := Globals.AllAirbases[abid]
		ab.Step(st)
	}

	// Step all aircrafts in the air
	for _, acid := range w.AircraftsInTheAir {
		ac := Globals.AllAircrafts[acid]
		ac.Step(st)
	}
}

func Step(st StepTime) {
	Globals.World.Step(st)
}
