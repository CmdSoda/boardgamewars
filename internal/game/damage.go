package game

import "github.com/CmdSoda/boardgamewars/internal/randomizer"

type DamageType int

const (
	DamageTypeFuselage DamageType = 0
	DamageTypeWing     DamageType = 1
	DamageTypeCockpit  DamageType = 2
	DamageTypeTurbine  DamageType = 3
	DamageTypeRudder   DamageType = 4
	DamageTypeFlaps    DamageType = 5
)

func (dt DamageType) String() string {
	switch dt {
	case DamageTypeFuselage:
		return "Fuselage"
	case DamageTypeWing:
		return "Wing"
	case DamageTypeCockpit:
		return "Cockpit"
	case DamageTypeTurbine:
		return "Turbine"
	case DamageTypeRudder:
		return "Rudder"
	case DamageTypeFlaps:
		return "Flaps"
	}
	return "Nothing"
}

func RollRandomDamage(weapondamage Hitpoints, targethp Hitpoints) []DamageType {
	damagecount := 1
	wdf := float32(weapondamage)
	thpf := float32(targethp)
	if wdf >= thpf * 0.8 {
		damagecount = 3
	} else if wdf >= thpf * 0.5 {
		damagecount = 2
	}
	var dtl []DamageType
	for i := 0; i < damagecount; i++ {
		r := randomizer.Roll1D10()
		switch r {
		case 1:
			dtl = append(dtl, DamageTypeFuselage)
		case 2:
			dtl = append(dtl, DamageTypeFuselage)
		case 3:
			dtl = append(dtl, DamageTypeFuselage)
		case 4:
			dtl = append(dtl, DamageTypeWing)
		case 5:
			dtl = append(dtl, DamageTypeWing)
		case 6:
			dtl = append(dtl, DamageTypeCockpit)
		case 7:
			dtl = append(dtl, DamageTypeTurbine)
		case 8:
			dtl = append(dtl, DamageTypeRudder)
		case 9:
			dtl = append(dtl, DamageTypeRudder)
		case 10:
			dtl = append(dtl, DamageTypeFlaps)
		}

	}
	return dtl
}
