package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/hexagon"
	"github.com/google/uuid"
	"github.com/looplab/fsm"
	"strings"
)

type AircraftId uuid.UUID

type AircraftIdList []AircraftId

func (al *AircraftIdList) PullFirst() AircraftId {
	id := (*al)[0]
	*al = append((*al)[1:])
	return id
}

type AircraftsMap map[AircraftId]*Aircraft

var currentAircraftShortId ShortId = 0

type Aircraft struct {
	AircraftId
	AircraftParametersId
	CountryName
	ShortId
	Altitude          AltitudeBand // Aktuelle Höhe.
	CurrentPosition   hexagon.HexPosition
	WeaponSystems     WeaponSystemList
	WeaponsConfigName string
	Damage            []DamageType // Eine Liste von Schäden
	Destroyed         bool
	Pilots            []PilotId
	StationedAt       AirbaseId
	StepsTaken        StepTime
	FSM               *fsm.FSM
	RepairTime        StepTime
	Waypoints         hexagon.PositionList
	CurrentWaypoint   int
	Destination       hexagon.HexPosition
	CalculatedPath    []hexagon.HexPosition
}

const (
	AcStateInTheAir      string = "in_the_air"
	AcStateParking       string = "parking"
	AcStateInDogfight    string = "in_dogfight"
	AcStateInMaintenance string = "in_maintenance"
	AcEventStart         string = "start"
	AcEventAttack        string = "attack"
	AcEventLand          string = "land"
	AcEventDisengage     string = "disengage"
	AcEventRepair        string = "repair"
	AcEventRepairDone    string = "repaired"
)

func (ac *Aircraft) GetHexPosition() hexagon.HexPosition {
	return hexagon.HexPosition{}
}

func (ac *Aircraft) GetHighestPilotRank() Code {
	highest := OF0
	for _, pid := range ac.Pilots {
		p := Globals.AllPilots[pid]
		if Code(p.Code) > highest {
			highest = Code(p.Code)
		}
	}
	return highest
}

func (ac *Aircraft) Repair() {
	ac.Damage = []DamageType{}
}

func (ac *Aircraft) ReviveAndRepair() {
	ac.Damage = []DamageType{}
	ac.Destroyed = false
}

func (ac *Aircraft) Rearm() {
	for i := range ac.WeaponSystems {
		ac.WeaponSystems[i].Depleted = false
	}
}

func (ac *Aircraft) AddPilot(id PilotId) {
	if len(ac.Pilots) >= ac.GetParameters().Seats {
		Log.Errorf("too many pilots in aircraft %d", ac.ShortId)
	}
	ac.Pilots = append(ac.Pilots, id)
}

func (ac *Aircraft) AssignToAB(id AirbaseId) bool {
	_, exist := Globals.AllAirbases[id]
	if exist {
		ac.StationedAt = id
		return true
	}
	return false
}

func NewAircraft(name string, weaponConfigName string, country CountryName) *Aircraft {
	ac := Aircraft{}
	ac.ShortId = currentAircraftShortId
	currentAircraftShortId = currentAircraftShortId + 1
	acpid, exist := GetAircraftParametersIdByName(name)
	if exist {
		ac.AircraftId = AircraftId(uuid.New())
		ac.AircraftParametersId = acpid
		ac.CountryName = country
		ac.WeaponSystems = CloneWeaponSystemList(acpid, weaponConfigName)
		ac.WeaponsConfigName = weaponConfigName
		for i := 0; i < len(ac.WeaponSystems); i++ {
			ac.WeaponSystems[i].InitWeaponSystem()
		}
		ac.Damage = make([]DamageType, 0)
		ac.FSM = fsm.NewFSM(AcStateParking, fsm.Events{
			{Name: AcEventStart, Src: []string{AcStateParking}, Dst: AcStateInTheAir},
			{Name: AcEventLand, Src: []string{AcStateInTheAir}, Dst: AcStateParking},
			{Name: AcEventAttack, Src: []string{AcStateInTheAir}, Dst: AcStateInDogfight},
			{Name: AcEventDisengage, Src: []string{AcStateInDogfight}, Dst: AcStateInTheAir},
			{Name: AcEventRepair, Src: []string{AcStateParking}, Dst: AcStateInMaintenance},
			{Name: AcEventRepairDone, Src: []string{AcStateInMaintenance}, Dst: AcStateParking},
		}, fsm.Callbacks{
			"enter_state": func(e *fsm.Event) { ac.enterState(e) },
		})
		Globals.AllAircrafts[ac.AircraftId] = &ac
		Log.Infof("new aircraft created: %s (AC%d)", name, ac.ShortId)
		AppendEvent(NewEvent(uuid.UUID(ac.AircraftId), SourceTypeAircraft, "created"))
		return &ac
	}
	return &ac
}

func (ac *Aircraft) enterState(e *fsm.Event) {
	switch e.Event {
	case AcEventRepair:
		ac.RepairTime = StepTime(20 * len(ac.Damage))
	}
	AppendEvent(NewEvent(uuid.UUID(ac.AircraftId), SourceTypeAircraft, e.Event))
}

func GetAircraftParametersIdByName(name string) (AircraftParametersId, bool) {
	for _, parameters := range Globals.AllAircraftParameters {
		if parameters.Name == name {
			return parameters.Id, true
		}
	}
	Log.Errorf("unkown aircraft name: %s", name)
	return InvalidAircraftParametersId, false
}

func (ac Aircraft) GetParameters() AircraftParameters {
	ap, _ := Globals.AllAircraftParameters[ac.AircraftParametersId]
	return ap
}

func (ac Aircraft) GetBestDogfightingWeapon() (WeaponSystem, bool) {
	var bestws WeaponSystem
	var max = 0
	exist := false
	if ac.WeaponSystems == nil {
		Log.Panicln("no weapon systems")
	}
	for _, system := range ac.WeaponSystems {
		if system.Depleted == false && system.Air2AirWeaponParameters != nil {
			if int(system.Air2AirWeaponParameters.Dogfighting) > max {
				bestws = system
				max = int(system.Dogfighting)
				exist = true
			}
		}
	}
	return bestws, exist
}

func (ac *Aircraft) AddLightDamage(dt DamageType) {
	ac.Damage = append(ac.Damage, dt)
}

func (ac *Aircraft) AddDamage(dtl []DamageType) {
	ac.Damage = append(ac.Damage, dtl...)
}

func (ac *Aircraft) DoDamageAssessment() {
	if len(ac.Damage) > ac.GetParameters().MaxDamagePoints {
		ac.Destroyed = true
	}

	// Falls die Kanzel getroffen wurde => Pilot tot?
}

func (ac *Aircraft) DoDamageWith(ws WeaponSystem) ([]DamageType, bool) {
	if ws.Air2AirWeaponParameters != nil {
		dhp := ws.Air2AirWeaponParameters.DoRandomDamage()
		acp := Globals.AllAircraftParameters[ac.AircraftParametersId]
		rd := RollRandomDamage(dhp, acp.MaxHitpoints)
		Globals.Statistics.WeaponPerformance.Damage(ws.Name, ac.AircraftId, len(rd))
		ac.AddDamage(rd)
		if len(ac.Damage) > acp.MaxDamagePoints {
			ac.Destroy()
			return rd, true
		}
	}
	return []DamageType{}, false
}

func (ac *Aircraft) DepleteWeapon(ws WeaponSystem) {
	if ws.Air2AirWeaponParameters != nil {
		for i, system := range ac.WeaponSystems {
			if system.Depleted == false && system.Air2AirWeaponParameters != nil &&
				ws.Air2AirWeaponParameters.Id == system.Air2AirWeaponParameters.Id {
				ac.WeaponSystems[i].Depleted = true
				return
			}
		}
	}
}

func (ac *Aircraft) FillSeatsWith(pl []PilotId) {
	if len(pl) > ac.GetParameters().Seats {
		panic("too many pilots")
	}
	for _, pilotid := range pl {
		ac.Pilots = append(ac.Pilots, pilotid)
	}

}

func (ac *Aircraft) FillSeatsWithNewPilots(nc Code) {
	pl := NewPilots(Globals.AllAircraftParameters[ac.AircraftParametersId].Seats, ac.CountryName, nc)
	ac.FillSeatsWith(pl)
}

func (ac *Aircraft) Destroy() {
	Log.Infof("AC%d destroyed", ac.ShortId)
	ac.Destroyed = true
}

//goland:noinspection GoUnhandledErrorResult
func (ac Aircraft) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "Aircraft(AC%d): %s\n", ac.ShortId, ac.GetParameters().Name)
	for _, pilotid := range ac.Pilots {
		p := Globals.AllPilots[pilotid]
		fmt.Fprintf(&b, "  Pilot: %s\n", p)
	}
	fmt.Fprint(&b, "  Damage: ")
	if len(ac.Damage) == 0 {
		fmt.Fprint(&b, "<no damage>")
	}
	for _, d := range ac.Damage {
		fmt.Fprintf(&b, d.String()+" ")
	}
	return b.String()
}

func (ac *Aircraft) IsDamaged() bool {
	return len(ac.Damage) > 0
}

func (ac *Aircraft) SetWaypoints(wps []hexagon.HexPosition) {
	ac.Waypoints = wps
	ac.CurrentWaypoint = 0
	ac.Destination = ac.Waypoints[0]
	ac.CalculatedPath = hexagon.CalculatePath(ac.CurrentPosition, ac.Destination)
}

func (ac *Aircraft) nextWaypoint() {
	ac.CurrentWaypoint++
	ac.Destination = ac.Waypoints[ac.CurrentWaypoint]
	ac.CalculatedPath = hexagon.CalculatePath(ac.CurrentPosition, ac.Destination)
}

func (ac *Aircraft) Step(st StepTime) {
	ac.StepsTaken = ac.StepsTaken + st
	switch ac.FSM.Current() {
	case AcStateInTheAir:
		ab := Globals.AllAirbases[ac.StationedAt]
		// Soll das Flugzeug landen?
		if ac.Destination == ac.CurrentPosition && ab.HexPosition == ac.CurrentPosition {
			err := ac.FSM.Event(AcEventLand)
			if err != nil {
				Log.Panicf("Unable to change AC%d to AcEventLand\n", ac.ShortId)
			}
			Globals.World.MoveAircraftInAirToAirbase(ac.AircraftId, ab.AirbaseId)
			ab.AddToParkingArea(ac.AircraftId)
		} else if ac.Destination != ac.CurrentPosition {
			// Soll das Flugzeug
			np := hexagon.GetNextPosition(ac.CurrentPosition, ac.CalculatedPath)
			ac.CurrentPosition = np
		} else if ac.Destination == ac.CurrentPosition {
			if ac.CurrentWaypoint < len(ac.Waypoints)-1 {
				ac.nextWaypoint()
			}
		}

	case AcStateParking:
		// TODO Soll das Aircraft sofort starten, wenn es Wegpunkte bekommen hat? Gibt es eine bessere Lösung?
		// Vielleicht will man Wegpunkte setzen können, ohne das Flugzeug zum Start zu zwingen.
		if len(ac.Waypoints) > 0 {
			err := ac.FSM.Event(AcEventStart)
			if err != nil {
				Log.Panicf("Unable to change AC%d to AcEventStart\n", ac.ShortId)
			}
		}
	case AcStateInDogfight:
	}
}
