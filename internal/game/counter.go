package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/hexagon"
)

type CounterType int

const (
	CounterTypeAircraft CounterType = 0
)

func (ct CounterType) String() string {
	switch ct {
	case CounterTypeAircraft:
		return "Aircraft"
	}
	return "Unknown"
}

type Counter struct {
	Position hexagon.HexPosition
	Type     CounterType
	Object   interface{}
}

func NewCounter(ct CounterType, obj interface{}) Counter {
	return Counter{
		Type:   ct,
		Object: obj,
	}
}

type CounterList []Counter

func (c Counter) String() string {
	return fmt.Sprintf("Type: %s, %s", c.Type.String(), c.Position.String())
}
