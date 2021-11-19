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

type Statistics struct {
	W2A2C WeaponNameVsAircraftParameterIdMap
	WVsA  *WinVsAircraftList
}

func NewStatistics() Statistics {
	s := Statistics{}
	s.W2A2C = map[string]map[AircraftParametersId]map[string]*WeaponStatistics{}
	s.WVsA = &WinVsAircraftList{}
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

func (w2a2c WeaponNameVsAircraftParameterIdMap) Dump() {
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
