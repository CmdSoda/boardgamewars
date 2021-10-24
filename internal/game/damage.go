package game

import "github.com/CmdSoda/boardgamewars/internal/randomizer"

type DamageType int

const (
	DamageTypeNothing  DamageType = -1
	DamageTypeFuselage DamageType = 0
	DamageTypeWing     DamageType = 1
	DamageTypeCockpit  DamageType = 2
	DamageTypeTurbine  DamageType = 3
	DamageTypeRudder   DamageType = 4
	DamageTypeFlaps    DamageType = 5
)

func RollRandomDamage() DamageType {
	r := randomizer.Roll1D10()
	switch r {
	case 1:
		return DamageTypeFuselage
	case 2:
		return DamageTypeFuselage
	case 3:
		return DamageTypeFuselage
	case 4:
		return DamageTypeWing
	case 5:
		return DamageTypeWing
	case 6:
		return DamageTypeCockpit
	case 7:
		return DamageTypeTurbine
	case 8:
		return DamageTypeRudder
	case 9:
		return DamageTypeRudder
	case 10:
		return DamageTypeFlaps
	}
	return DamageTypeNothing
}
