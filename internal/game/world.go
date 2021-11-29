package game

import "fmt"

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

func (w *World) MoveAircraftToAir(acid AircraftId) {

}

func (w *World) MoveAircraftToAirbase(acid AircraftId) {

}

func (w *World) Step(st StepTime) {
	w.CurrentStep = w.CurrentStep + st
	fmt.Println(st)
	for _, abid := range w.Airbases {
		ab := Globals.AllAirbases[abid]
		ab.Step(st)
	}
}

func Step(st StepTime) {
	Globals.World.Step(st)
}
