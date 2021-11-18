package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/nato"
	"github.com/google/uuid"
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
	WarPartyId
	ShortId
	Altitude           AltitudeBand // Aktuelle Höhe.
	CurrentPosition    FloatPosition
	NextTargetLocation FloatPosition // Das ist die FloatPosition, die das Flugzeug jetzt ansteuert.
	WeaponSystems      WeaponSystemList
	WeaponsConfigName  string
	Damage             []DamageType // Eine Liste von Schäden
	Destroyed          bool
	Pilots             []PilotId
	StationedAt        AirbaseId
}

func (a Aircraft) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "Aircraft(AC%d): %s\n", a.ShortId, a.GetParameters().Name)
	for _, pilotid := range a.Pilots {
		p := Globals.AllPilots[pilotid]
		fmt.Fprintf(&b, "  Pilot: %s\n", p)
	}
	fmt.Fprint(&b, "  Damage: ")
	if len(a.Damage) == 0 {
		fmt.Fprint(&b, "<no damage>")
	}
	for _, d := range a.Damage {
		fmt.Fprintf(&b, d.String()+" ")
	}
	return b.String()
}

func (a *Aircraft) Repair() {
	a.Damage = []DamageType{}
}

func (a *Aircraft) ReviveAndRepair() {
	a.Damage = []DamageType{}
	a.Destroyed = false
}

func (a *Aircraft) Rearm() {
	for i, _ := range a.WeaponSystems {
		a.WeaponSystems[i].Depleted = false
	}
}

func (a *Aircraft) AddPilot(id PilotId) {
	if len(a.Pilots) >= a.GetParameters().Seats {
		Log.Errorf("too many pilots in aircraft %d", a.ShortId)
	}
	a.Pilots = append(a.Pilots, id)
}

func (a *Aircraft) AssignToAB(id AirbaseId) bool {
	_, exist := Globals.AllAirbases[id]
	if exist {
		a.StationedAt = id
		return true
	}
	return false
}

func NewAircraft(name string, weaponConfigName string, warpartyid WarPartyId) *Aircraft {
	ac := Aircraft{}
	ac.ShortId = currentAircraftShortId
	currentAircraftShortId = currentAircraftShortId + 1
	acpid, exist := GetAircraftParametersIdByName(name)
	if exist {
		ac.AircraftId = AircraftId(uuid.New())
		ac.AircraftParametersId = acpid
		ac.WarPartyId = warpartyid
		ac.WeaponSystems = CloneWeaponSystemList(acpid, weaponConfigName)
		ac.WeaponsConfigName = weaponConfigName
		for i := 0; i < len(ac.WeaponSystems); i++ {
			ac.WeaponSystems[i].InitWeaponSystem()
		}
		ac.Damage = make([]DamageType, 0)
		Globals.AllAircrafts[ac.AircraftId] = &ac
		Log.Infof("new aircraft created: %s (%d)", name, ac.ShortId)
		return &ac
	}
	return &ac
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

func (a Aircraft) GetParameters() AircraftParameters {
	ap, _ := Globals.AllAircraftParameters[a.AircraftParametersId]
	return ap
}

func (a Aircraft) GetBestDogfightingWeapon() (WeaponSystem, bool) {
	var bestws WeaponSystem
	var max = 0
	exist := false
	if a.WeaponSystems == nil {
		panic("no weapon systems")
	}
	for _, system := range a.WeaponSystems {
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

func (a *Aircraft) AddLightDamage(dt DamageType) {
	a.Damage = append(a.Damage, dt)
}

func (a *Aircraft) DoDamageAssessment() {
	if len(a.Damage) > a.GetParameters().MaxDamagePoints {
		a.Destroyed = true
	}

	// Falls die Kanzel getroffen wurde => Pilot tot?
}

func (a *Aircraft) DoDamageWith(ws WeaponSystem) DamageType {
	if ws.Air2AirWeaponParameters != nil {
		dhp := ws.Air2AirWeaponParameters.DoRandomDamage()
		if dhp <= a.GetParameters().MaxHitpoints {
			rd := RollRandomDamage()
			a.AddLightDamage(rd)
			return rd
		}
	}
	return DamageTypeNothing
}

func (a *Aircraft) DepleteWeapon(ws WeaponSystem) {
	if ws.Air2AirWeaponParameters != nil {
		for i, system := range a.WeaponSystems {
			if system.Depleted == false && system.Air2AirWeaponParameters != nil &&
				ws.Air2AirWeaponParameters.Id == system.Air2AirWeaponParameters.Id {
				a.WeaponSystems[i].Depleted = true
				return
			}
		}
	}
}

func (a *Aircraft) FillSeatsWith(pl []PilotId) {
	if len(pl) > a.GetParameters().Seats {
		panic("too many pilots")
	}
	for _, pilotid := range pl {
		a.Pilots = append(a.Pilots, pilotid)
	}

}

func (a *Aircraft) FillSeatsWithNewPilots(nc nato.Code) {
	pl := NewPilots(Globals.AllAircraftParameters[a.AircraftParametersId].Seats, a.WarPartyId, nc)
	a.FillSeatsWith(pl)
}

func (a *Aircraft) Destroy() {
	Log.Infof("AC%d destroyed", a.ShortId)
	a.Destroyed = true
}
