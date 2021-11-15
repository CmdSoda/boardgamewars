package game

import (
	"fmt"
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

type AircraftsMap map[AircraftId]Aircraft

type ShortId int
var currentShortId ShortId = 0

type Aircraft struct {
	AircraftId
	AircraftParametersId
	WarPartyId
	ShortId
	Altitude           AltitudeBand // Aktuelle Höhe.
	CurrentPosition    FloatPosition
	NextTargetLocation FloatPosition // Das ist die FloatPosition, die das Flugzeug jetzt ansteuert.
	WeaponSystems      WeaponSystemList
	Damage             []DamageType // Eine Liste von Schäden
	Destroyed          bool
	Pilots             []PilotId
	StationedAt        AirbaseId
}

func (a Aircraft) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "Aircraft %d\n", a.ShortId)
	fmt.Fprintf(&b, "%s\nPilots: ", a.GetParameters().Name)
	for _, pilotid := range a.Pilots {
		p := Globals.AllPilots[pilotid]
		fmt.Fprintf(&b, p.String()+" ")
	}
	fmt.Fprint(&b, "\nDamage: ")
	for _, d := range a.Damage {
		fmt.Fprintf(&b, d.String()+" ")
	}
	if len(a.Damage) == 0 {
		fmt.Fprint(&b, "\n")
	}
	return b.String()
}

func (a *Aircraft) AddPilot(id PilotId) {
	if len(a.Pilots) >= a.GetParameters().Seats {
		panic("too many pilots in aircraft")
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

func NewAircraft(name string, configurationName string, warpartyid WarPartyId) Aircraft {
	ac := Aircraft{}
	ac.ShortId = currentShortId
	currentShortId = currentShortId + 1
	acpid, exist := GetAircraftParametersIdByName(name)
	if exist {
		ac.AircraftId = AircraftId(uuid.New())
		ac.AircraftParametersId = acpid
		ac.WarPartyId = warpartyid
		ac.WeaponSystems = GetWeaponSystemList(acpid, configurationName)
		for i := 0; i < len(ac.WeaponSystems); i++ {
			ac.WeaponSystems[i].InitWeaponSystem()
		}
		ac.Damage = make([]DamageType, 0)
		Globals.AllAircrafts[ac.AircraftId] = ac
		return ac
	}
	return ac
}

func GetAircraftParametersIdByName(name string) (AircraftParametersId, bool) {
	for _, parameters := range Globals.AllAircraftParameters {
		if parameters.Name == name {
			return parameters.Id, true
		}
	}
	panic("unkown aircraft name: " + name)
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
