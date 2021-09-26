package world

type HillType uint

//goland:noinspection ALL
const (
	// Very plain, like a dream.
	HillPlain HillType = iota
	// It goes up and down just a little.
	HillRough
	// Very high Slope.
	HillSteep
	// Like a wall.
	HillFace
)

type GroundType uint

//goland:noinspection ALL
const (
	GroundStreet GroundType = iota
	GroundGras
	GroundMud
	GroundWood
	GroundUrban
	GroundWaterShallow
	GroundWaterDeep
)
