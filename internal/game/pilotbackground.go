package game

import "github.com/CmdSoda/boardgamewars/internal/countrycodes"

type PilotBackground struct {
	Age     int
	Country countrycodes.Code
	Born    string
}
