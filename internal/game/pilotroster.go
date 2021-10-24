package game

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

type PilotRoster map[uuid.UUID]Pilot

func (pr PilotRoster) Add(p Pilot) bool {
	if pr.Exist(p.UUID) {
		return false
	}
	pr[p.UUID] = p
	return true
}

func (pr PilotRoster) Exist(id uuid.UUID) bool {
	_, exist := pr[id]
	return exist
}

func (pr PilotRoster) GetPilot(id uuid.UUID) Pilot {
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

var TheRoster PilotRoster

func NewPilotRoster() {
	TheRoster = PilotRoster{}
}
