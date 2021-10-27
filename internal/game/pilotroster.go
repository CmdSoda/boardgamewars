package game

import (
	"fmt"
	"strings"
)

type PilotRoster map[PilotId]Pilot

func (pr PilotRoster) Add(p Pilot) bool {
	_, exist := Globals.PilotRoster[p.PilotId]
	if exist {
		return false
	}
	pr[p.PilotId] = p
	return true
}

func (pr PilotRoster) GetPilot(id PilotId) Pilot {
	return pr[id]
}

func (pr PilotRoster) String() string {
	var b strings.Builder
	fmt.Fprint(&b, "Roster:\n")
	for _, pilot := range pr {
		fmt.Fprintf(&b, "%s\n", pilot.String())
	}
	return b.String()
}

func NewPilotRoster() {
	Globals.PilotRoster = PilotRoster{}
}
