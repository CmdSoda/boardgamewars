package game

import "fmt"

type FloatPosition struct {
	X float64
	Y float64
}

func (p FloatPosition) String() string {
	return fmt.Sprintf("%f, %f", p.X, p.Y)
}
