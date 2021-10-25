// Royal Air Force
// https://en.wikipedia.org/wiki/Ranks_and_insignia_of_NATO_air_forces_officers

package military

import "github.com/CmdSoda/boardgamewars/internal/nato"

// PilotRanksUK Royal Air Force
var PilotRanksUK = []FlightRank{
	{
		Name:  "Pilot Officer",
		Short: "Plt Off",
		Code:  nato.OF1,
	},
	{
		Name:  "Flying Officer",
		Short: "Fg Off",
		Code:  nato.OF1,
	},
	{
		Name:  "Flight Lieutenant",
		Short: "Flt Lt",
		Code:  nato.OF2,
	},
	{
		Name:  "Squadron Leader",
		Short: "Sqn Ldr",
		Code:  nato.OF3,
	},
	{
		Name:  "Wing Commander",
		Short: "Wg Cdr",
		Code:  nato.OF4,
	},
	{
		Name:  "Group Captain",
		Short: "Gp Cpt",
		Code:  nato.OF5,
	},
}
