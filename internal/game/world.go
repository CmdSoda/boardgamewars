package game

import "fmt"

type World struct {
	CurrentStep StepTime
}

func NewWorld() World {
	return World{}
}

func (w *World) Step(st StepTime) {
	w.CurrentStep = w.CurrentStep + st
	fmt.Println(st)
}

func Step(st StepTime) {
	Globals.World.Step(st)
}
