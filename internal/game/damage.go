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

func GetDamageTypeFromString(dts string) DamageType {
	switch dts {
	case "Fuselage":
		return DamageTypeFuselage
	case "Wing":
		return DamageTypeWing
	case "Cockpit":
		return DamageTypeCockpit
	case "Turbine":
		return DamageTypeTurbine
	case "Rudder":
		return DamageTypeRudder
	case "Flaps":
		return DamageTypeFlaps
	}
	return DamageTypeFuselage
}

func RollRandomDamage(weapondamage Hitpoints, targethp Hitpoints) []DamageType {
	damagecount := 0
	wdf := float32(weapondamage)
	thpf := float32(targethp)
	if wdf >= thpf*Globals.Settings.DamageCountStage1Percent/100.0 {
		damagecount = Globals.Settings.DamageCountStage1Value
	} else if wdf >= thpf*Globals.Settings.DamageCountStage2Percent/100.0 {
		damagecount = Globals.Settings.DamageCountStage2Value
	} else if wdf >= thpf*Globals.Settings.DamageCountStage3Percent/100.0 {
		damagecount = Globals.Settings.DamageCountStage3Value
	}
	var dtl []DamageType
	for i := 0; i < damagecount; i++ {
		r := randomizer.Roll1DN(len(Globals.Settings.RandomDamage))
		dtl = append(dtl, Globals.Settings.RandomDamage[r-1])
	}
	return dtl
}
