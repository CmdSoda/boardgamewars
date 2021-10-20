package vector

//
// Vector 2
//

type Vector struct {
	X float64
	Y float64
}

//goland:noinspection ALL
func NewVector(x float64, y float64) Vector {
	return Vector{
		X: x,
		Y: y,
	}
}

func (v Vector) Minus(v2 Vector) Vector {
	return Vector{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}
