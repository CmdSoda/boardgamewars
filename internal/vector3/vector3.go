package vector3

import "math"

//
// Vector 3
//

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

func NewVector(x float64, y float64, z float64) Vector3 {
	return Vector3{
		X: x,
		Y: y,
		Z: z,
	}
}

func (v Vector3) Sub(v2 Vector3) Vector3 {
	return Vector3{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
		Z: v.Z - v2.Z,
	}
}

func (v Vector3) Cross(ov Vector3) Vector3 {
	return Vector3{
		v.Y*ov.Z - v.Z*ov.Y,
		v.Z*ov.X - v.X*ov.Z,
		v.X*ov.Y - v.Y*ov.X,
	}
}

func (v Vector3) Distance(ov Vector3) float64 { return v.Sub(ov).Norm() }

func (v Vector3) Norm() float64 { return math.Sqrt(v.Dot(v)) }

func (v Vector3) Dot(ov Vector3) float64 { return v.X*ov.X + v.Y*ov.Y + v.Z*ov.Z }
