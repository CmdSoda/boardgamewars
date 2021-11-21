package game

type StepTime int

type Stepper interface {
	Step(st StepTime)
}
