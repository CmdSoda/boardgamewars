package game

import "fmt"

type World struct {
}

func NewWorld() World {
	return World{}
}

func (w World) Step(st StepTime) {
	fmt.Println(st)
}
