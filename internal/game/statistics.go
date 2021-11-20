package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/nato"
)

type Statistics struct {
	WeaponPerformance  WeaponPerformanceMap
	AircraftVsAircraft AircraftVersusPerformanceList
}

// WeaponPerformanceMap ist eine Map auf Waffensystem-Name auf AircraftParametersId auf Config-Name
// Beispiel:
// m := WeaponPerformanceMap{}
// m["Aim-9"][acid]["Default"].Hit += 1
type WeaponPerformanceMap map[string]map[AircraftParametersId]map[string]*WeaponPerformanceStatistics

type WeaponPerformanceStatistics struct {
	Hit     int
	NotHit  int
	Damage1 int
	Damage2 int
	Damage3 int
}

// AircraftVersusPerformanceList speichert Win-Statistiken:
// m[acF14]["Default"][acMig29]["Default"] = WinStatistics{ Won: 23, Lost: 67, Draw: 6 }
type AircraftVersusPerformanceList []*WinStatistics

type WinAircraftParameters struct {
	AircraftParametersId
	ConfigName string
	PilotRank  nato.Code
}

type WinAircraftStats struct {
	Won            int
	Draw           int
	Lost           int
	AdvantageCount int
}

type WinType int

const (
	WinTypeDraw WinType = 0
	WinTypeWon  WinType = 1
)

type WinStatistics struct {
	AC1Params WinAircraftParameters
	AC2Params WinAircraftParameters
	AC1Stats  WinAircraftStats
	AC2Stats  WinAircraftStats
}

type WeaponDmgCountMap map[string]map[AircraftParametersId]*DamageStatistics

type DamageStatistics struct {
	Damage1 int
	Damage2 int
	Damage3 int
}

func NewStatistics() Statistics {
	s := Statistics{}
	s.WeaponPerformance = map[string]map[AircraftParametersId]map[string]*WeaponPerformanceStatistics{}
	s.AircraftVsAircraft = AircraftVersusPerformanceList{}
	return s
}

func (w *AircraftVersusPerformanceList) Win(acid1 AircraftId, acid2 AircraftId, wtype WinType) {

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
				(*w)[i].AC1Stats.Draw = (*w)[i].AC1Stats.Draw + 1
				(*w)[i].AC2Stats.Draw = (*w)[i].AC2Stats.Draw + 1
			case WinTypeWon:
				(*w)[i].AC1Stats.Won = (*w)[i].AC1Stats.Won + 1
				(*w)[i].AC2Stats.Lost = (*w)[i].AC2Stats.Lost + 1
			}
			return
		} else if (*w)[i].AC1Params == wap2 && (*w)[i].AC2Params == wap1 {
			switch wtype {
			case WinTypeDraw:
				(*w)[i].AC1Stats.Draw = (*w)[i].AC1Stats.Draw + 1
				(*w)[i].AC2Stats.Draw = (*w)[i].AC2Stats.Draw + 1
			case WinTypeWon:
				(*w)[i].AC2Stats.Won = (*w)[i].AC2Stats.Won + 1
				(*w)[i].AC1Stats.Lost = (*w)[i].AC1Stats.Lost + 1
			}
			return
		}
	}

	ws := WinStatistics{
		AC1Params: wap1,
		AC2Params: wap2,
		AC1Stats: WinAircraftStats{
			Won:            0,
			Draw:           0,
			Lost:           0,
			AdvantageCount: 0,
		},
		AC2Stats: WinAircraftStats{
			Won:            0,
			Draw:           0,
			Lost:           0,
			AdvantageCount: 0,
		},
	}

	switch wtype {
	case WinTypeDraw:
		ws.AC1Stats = WinAircraftStats{
			Won:            0,
			Draw:           1,
			Lost:           0,
			AdvantageCount: 0,
		}
		ws.AC2Stats = WinAircraftStats{
			Won:            0,
			Draw:           1,
			Lost:           0,
			AdvantageCount: 0,
		}
	case WinTypeWon:
		ws.AC1Stats = WinAircraftStats{
			Won:            1,
			Draw:           0,
			Lost:           0,
			AdvantageCount: 0,
		}
		ws.AC2Stats = WinAircraftStats{
			Won:            0,
			Draw:           0,
			Lost:           1,
			AdvantageCount: 0,
		}
	}

	// Diese Kombination gibt es noch nicht und es muss eine erstellt werden.
	*w = append(*w, &ws)
}

func (w *AircraftVersusPerformanceList) Dump() {
	fmt.Println("Statistics: Aircraft vs Aircraft")
	for _, wstat := range *w {
		acp1 := Globals.AllAircraftParameters[wstat.AC1Params.AircraftParametersId]
		acp2 := Globals.AllAircraftParameters[wstat.AC2Params.AircraftParametersId]
		var acwr1 float32
		var acwr2 float32
		if wstat.AC1Stats.Won+wstat.AC1Stats.Won == 0 {
			acwr1 = 0
		} else {
			acwr1 = float32(wstat.AC1Stats.Won) / float32(wstat.AC1Stats.Won+wstat.AC2Stats.Won) * 100
		}
		if wstat.AC1Stats.Won+wstat.AC2Stats.Won == 0 {
			acwr2 = 0
		} else {
			acwr2 = float32(wstat.AC2Stats.Won) / float32(wstat.AC1Stats.Won+wstat.AC2Stats.Won) * 100
		}
		samples := wstat.AC1Stats.Won + wstat.AC1Stats.Lost + wstat.AC1Stats.Draw
		fmt.Printf("%s(%.1f%%) vs %s(%.1f%%) (%d samples)\n", acp1.Name, acwr1, acp2.Name, acwr2, samples)
	}
}

func (w2a2c WeaponPerformanceMap) Dump() {
	fmt.Println("Statistics: Weapon Performance")
	for wname := range w2a2c {
		fmt.Println(wname)
		for acid := range w2a2c[wname] {
			acp := Globals.AllAircraftParameters[acid]
			fmt.Println("  " + acp.Name)
			for config := range w2a2c[wname][acid] {
				stats := w2a2c[wname][acid][config]
				count := stats.Damage1 + stats.Damage2 + stats.Damage3
				average := float32(stats.Damage1+stats.Damage2*2+stats.Damage3*3) /
					float32(count)
				sum := w2a2c[wname][acid][config].Hit + w2a2c[wname][acid][config].NotHit
				hitpc := float32(w2a2c[wname][acid][config].Hit) / float32(sum) * 100
				fmt.Printf("    %s Hit=%.1f%%, DmgAvg=%.1f (%d samples)\n", config, hitpc, average, sum)
			}
		}
	}
}

func (w2a2c *WeaponPerformanceMap) createMaps(weaponName string, acid AircraftId) {
	ac := Globals.AllAircrafts[acid]

	_, wnexist := (*w2a2c)[weaponName]
	if !wnexist {
		(*w2a2c)[weaponName] = map[AircraftParametersId]map[string]*WeaponPerformanceStatistics{}
	}

	_, apidexist := (*w2a2c)[weaponName][ac.AircraftParametersId]
	if !apidexist {
		(*w2a2c)[weaponName][ac.AircraftParametersId] = map[string]*WeaponPerformanceStatistics{}
	}

	_, cexist := (*w2a2c)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName]
	if !cexist {
		(*w2a2c)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName] = &WeaponPerformanceStatistics{}
	}
}

func (w2a2c *WeaponPerformanceMap) Hit(weaponName string, acid AircraftId) {
	w2a2c.createMaps(weaponName, acid)
	ac := Globals.AllAircrafts[acid]
	// Um 1 erhöhen.
	(*w2a2c)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].Hit =
		(*w2a2c)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].Hit + 1
}

func (w2a2c *WeaponPerformanceMap) NotHit(weaponName string, acid AircraftId) {
	w2a2c.createMaps(weaponName, acid)
	ac := Globals.AllAircrafts[acid]
	// Um 1 erhöhen.
	(*w2a2c)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].NotHit =
		(*w2a2c)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].NotHit + 1
}

func (w2a2c *WeaponPerformanceMap) Damage(weaponName string, acid AircraftId, dmg int) {
	w2a2c.createMaps(weaponName, acid)
	ac := Globals.AllAircrafts[acid]

	switch dmg {
	case 1:
		(*w2a2c)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].Damage1 =
			(*w2a2c)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].Damage1 + 1
	case 2:
		(*w2a2c)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].Damage2 =
			(*w2a2c)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].Damage2 + 1
	case 3:
		(*w2a2c)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].Damage3 =
			(*w2a2c)[weaponName][ac.AircraftParametersId][ac.WeaponsConfigName].Damage3 + 1
	}

}