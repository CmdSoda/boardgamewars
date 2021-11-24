package game

import "fmt"

type World struct {
	CurrentStep StepTime
	Machines    []*Stepper
}

func NewWorld() World {
	return World{Machines: []*Stepper{}}
}

func (w *World) Add(st *Stepper) {
	//var s Stepper = st
	w.Machines = append(w.Machines, st)
}

func (w *World) Step(st StepTime) {
	w.CurrentStep = w.CurrentStep + st
	fmt.Println(st)
	for _, machine := range w.Machines {
		(*machine).Step(st)
	}
}

func Step(st StepTime) {
	Globals.World.Step(st)
}
