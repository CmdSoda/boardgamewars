package game

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

type SourceType int

const (
	SourceTypeUnknown  SourceType = 0
	SourceTypeAircraft SourceType = 1
	SourceTypePilot    SourceType = 2
	SourceTypeAirbase  SourceType = 3
)

func (st SourceType) String() string {
	switch st {
	case SourceTypeUnknown:
		return "Unknown"
	case SourceTypeAircraft:
		return "Aircraft"
	case SourceTypePilot:
		return "Pilot"
	case SourceTypeAirbase:
		return "Airbase"
	default:
		return "Unknown"
	}
}

type Event struct {
	Source uuid.UUID
	SourceType
	StepTime
	Message string
}

func NewEvent(thisId uuid.UUID, st SourceType, msg string) Event {
	return Event{
		Source:     thisId,
		SourceType: st,
		StepTime:   Globals.World.CurrentStep,
		Message:    msg,
	}
}

func (e Event) String() string {
	return fmt.Sprintf("%s %s: %s", e.SourceType.String(), e.Source.String(), e.Message)
}

type EventList []Event

func NewEventList() EventList {
	return make(EventList, 0)
}

//goland:noinspection GoUnhandledErrorResult
func (el *EventList) String() string {
	var sb strings.Builder
	for _, event := range *el {
		fmt.Fprintf(&sb, "%s\n", event.String())
	}
	return sb.String()
}

func AppendEvent(e Event) {
	Globals.EventList = append(Globals.EventList, e)
}
