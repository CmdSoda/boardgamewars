package world

type HillType uint

//goland:noinspection ALL
const (
	HillPlain HillType = iota
	HillRough
	HillSteep
)

type GroundType uint

//goland:noinspection ALL
const (
	GroundStreet GroundType = iota
	GroundGras
	GroundMud
	GroundWood
	GroundUrban
	GroundWater
)
