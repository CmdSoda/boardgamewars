package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
)

type FlightRank struct {
	Name  string
	Short string
	Code  int
}

func (f FlightRank) String() string {
	if f.Short != "" {
		return fmt.Sprintf("%s", f.Short)
	} else {
		return f.Name
	}
}

type FlightRankList []FlightRank

func (frl FlightRankList) RandomPick() FlightRank {
	roll := randomizer.Roll1DN(len(frl))
	return frl[roll-1]
}

func GetAllRanksWithNatoCode(list []FlightRank, noc nato.Code) FlightRankList {
	frl := make(FlightRankList, 0)
	for _, rank := range list {
		if nato.Code(rank.Code) == noc {
			frl = append(frl, rank)
		}
	}
	return frl
}

func GetFlightRank(list []FlightRank, noc nato.Code) FlightRank {
	frl := GetAllRanksWithNatoCode(list, noc)
	return frl.RandomPick()
}

func NewRank(cn CountryName, noc nato.Code) FlightRank {
	return GetFlightRank(Globals.CountryDataMap[cn].FlightRankList, noc)
}
