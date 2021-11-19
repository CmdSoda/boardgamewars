package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/nato"
)

// WeaponNameVsAircraftParameterIdMap ist eine Map auf Waffensystem-Name auf AircraftParametersId auf Config-Name
// Beispiel:
// m := WeaponNameVsAircraftParameterIdMap{}
// m["Aim-9"][acid]["Default"].Hit += 1
type WeaponNameVsAircraftParameterIdMap map[string]map[AircraftParametersId]map[string]*WeaponStatistics

type WeaponStatistics struct {
	Hit    int
	NotHit int
}

// WinVsAircraftList speichert Win-Statistiken:
// m[acF14]["Default"][acMig29]["Default"] = WinStatistics{ Won: 23, Lost: 67, Draw: 6 }
type WinVsAircraftList []*WinStatistics

type WinAircraftParameters struct {
	AircraftParametersId
	ConfigName string
	PilotRank  nato.Code
}

type WinType int

const (
	WinTypeDraw WinType = 0
	WinTypeWon  WinType = 1
)

type WinStatistics struct {
	AC1Params WinAircraftParameters
	AC2Params WinAircraftParameters
	AC1Won    int
	AC2Won    int
	Draw      int
}

type DamageMap map[string]map[AircraftParametersId]*DamageStatistics

type DamageStatistics struct {
	Damage1 int
	Damage2 int
	Damage3 int
}

func (dm *DamageMap) Add(weaponname string, apid AircraftParametersId, dmg int) {
	if _, exist := (*dm)[weaponname]; !exist {
		(*dm)[weaponname] = map[AircraftParametersId]*DamageStatistics{}
	}
	if _, exist := (*dm)[weaponname][apid]; !exist {
		(*dm)[weaponname][apid] = &DamageStatistics{0, 0, 0}
	}
	switch dmg {
	case 1:
		(*dm)[weaponname][apid].Damage1 = (*dm)[weaponname][apid].Damage1 + 1
	case 2:
		(*dm)[weaponname][apid].Damage2 = (*dm)[weaponname][apid].Damage2 + 1
	case 3:
		(*dm)[weaponname][apid].Damage3 = (*dm)[weaponname][apid].Damage3 + 1
	}
}

func (dm *DamageMap) Dump() {
	fmt.Println("Statistics: Weapon vs Aircraft Damage Count")
	for weaponname, _ := range *dm {
		fmt.Println(weaponname)
		for apid, _ := range (*dm)[weaponname] {
			acp := Globals.AllAircraftParameters[apid]
			stats := (*dm)[weaponname][apid]
			count := stats.Damage1 + stats.Damage2 + stats.Damage3
			average := float32(stats.Damage1 + stats.Damage2 * 2 + stats.Damage3 * 3)/
				float32(count)
			fmt.Printf("  %s = %.2f (%d samples)\n", acp.Name, average, count)
		}
	}
}

type Statistics struct {
	W2A2C  WeaponNameVsAircraftParameterIdMap
	WVsA   WinVsAircraftList
	DmgVsA DamageMap
}

func NewStatistics() Statistics {
	s := Statistics{}
	s.W2A2C = map[string]map[AircraftParametersId]map[string]*WeaponStatistics{}
	s.WVsA = WinVsAircraftList{}
	s.DmgVsA = map[string]map[AircraftParametersId]*DamageStatistics{}
	return s
}

func (w *WinVsAircraftList) Win(acid1 AircraftId, acid2 AircraftId, wtype WinType) {

	ac1 := Globals.AllAircrafts[acid1]
	ac2 := Globals.AllAircrafts[acid2]

	wap1 := WinAircraftParameters{
		AircraftParametersId: ac1.AircraftParametersId,
		ConfigName:           ac1.WeaponsConfigName,
		PilotRank:            ac1.GetHighestPilotRank(),
	}

	wap2 := WinAircraftParameters{
		AircraftParametersId: ac2.AircraftParametersId,
		ConfigName:           ac2.WeaponsConfigName,
		PilotRank:            ac2.GetHighestPilotRank(),
	}

	// Die Möglichkeit besteht, dass es den Eintrag schon gibt. Suchen...
	for i, _ := range *w {
		if (*w)[i].AC1Params == wap1 && (*w)[i].AC2Params == wap2 {
			switch wtype {
			case WinTypeDraw:
				(*w)[i].Draw = (*w)[i].Draw + 1
			case WinTypeWon:
				(*w)[i].AC1Won = (*w)[i].AC1Won + 1
			}
			return
		} else if (*w)[i].AC1Params == wap2 && (*w)[i].AC2Params == wap1 {
			switch wtype {
			case WinTypeDraw:
				(*w)[i].Draw = (*w)[i].Draw + 1
			case WinTypeWon:
				(*w)[i].AC2Won = (*w)[i].AC2Won + 1
			}
			return
		}
	}

	ws := WinStatistics{
		AC1Params: wap1,
		AC2Params: wap2,
		AC1Won:    0,
		AC2Won:    0,
		Draw:      0,
	}

	switch wtype {
	case WinTypeDraw:
		ws.Draw = 1
	case WinTypeWon:
		ws.AC1Won = 1
	}

	// Diese Kombination gibt es noch nicht und es muss eine erstellt werden.
	*w = append(*w, &ws)
}

func (w *WinVsAircraftList) Dump() {
	fmt.Println("Statistics: Aircraft vs Aircraft")
	for _, wstat := range *w {
		acp1 := Globals.AllAircraftParameters[wstat.AC1Params.AircraftParametersId]
		acp2 := Globals.AllAircraftParameters[wstat.AC2Params.AircraftParametersId]
		var acwr1 float32
		var acwr2 float32
		if wstat.AC1Won+wstat.AC1Won == 0 {
			acwr1 = 0
		} else {
			acwr1 = float32(wstat.AC1Won) / float32(wstat.AC1Won+wstat.AC2Won) * 100
		}
		if wstat.AC1Won+wstat.AC2Won == 0 {
			acwr2 = 0
		} else {
			acwr2 = float32(wstat.AC2Won) / float32(wstat.AC1Won+wstat.AC2Won) * 100
		}
		samples := wstat.AC1Won + wstat.AC2Won + wstat.Draw
		fmt.Printf("%s(%.1f%%) vs %s(%.1f%%) (%d samples)\n", acp1.Name, acwr1, acp2.Name, acwr2, samples)
	}
}

func (w2a2c WeaponNameVsAircraftParameterIdMap) Dump() {
	fmt.Println("Statistics: Weapon vs Aircraft Hit%")
	for wname := range w2a2c {
		fmt.Println(wname)
		for acid := range w2a2c[wname] {
			acp := Globals.AllAircraftParameters[acid]
			fmt.Println("  " + acp.Name)
			for config := range w2a2c[wname][acid] {
				sum := w2a2c[wname][acid][config].Hit + w2a2c[wname][acid][config].NotHit
				hitpc := float32(w2a2c[wname][acid][config].Hit) / float32(sum) * 100
				fmt.Printf("    %s Hit %.1f%% (%d samples)\n", config, hitpc, sum)
			}
		}
	}
}

func (w2a2c *WeaponNameVsAircraftParameterIdMap) createMaps(weaponName string, acid AircraftId) {
	ac := Globals.AllAircrafts[acid]

	_, wnexist := (*w2a2c)[weaponName]
	if !wnexist {
		(*w2a2c)[weaponName] = map[AircraftParametersId]map[string]*WeaponStatistics{}
	}

	_, apidexist := (*w2a2c)[weaponName][ac.AircraftParametersId]
	if !apidexist {
		(*w2a2c)[weaponName][ac.AircraftParametersId] = map[string]*WeaponStatistics{}
	}

	_, cexist := (*w2a2c)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName]
	if !cexist {
		(*w2a2c)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName] = &WeaponStatistics{}
	}
}

func (w2a2c *WeaponNameVsAircraftParameterIdMap) Hit(weaponName string, acid AircraftId) {
	w2a2c.createMaps(weaponName, acid)
	ac := Globals.AllAircrafts[acid]
	// Um 1 erhöhen.
	(*w2a2c)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].Hit =
		(*w2a2c)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].Hit + 1
}

func (w2a2c *WeaponNameVsAircraftParameterIdMap) NotHit(weaponName string, acid AircraftId) {
	w2a2c.createMaps(weaponName, acid)
	ac := Globals.AllAircrafts[acid]
	// Um 1 erhöhen.
	(*w2a2c)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].NotHit =
		(*w2a2c)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].NotHit + 1
}
