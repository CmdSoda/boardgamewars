package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
	"strings"
)

type DogfightPosition int

const (
	DogfightPositionBehindEnemiesTailOptimal DogfightPosition = 3
	DogfightPositionBehindEnemiesTail        DogfightPosition = 2
	DogfightPositionAdventage                DogfightPosition = 1
	DogfightPositionTossup                   DogfightPosition = 0
	DogfightPositionIgnore                   DogfightPosition = 0
	DogfightPositionDisadvantage             DogfightPosition = -1
	DogfightPositionEnemyAtMySix             DogfightPosition = -2
	DogfightPositionEnemyAtMySixOptimal      DogfightPosition = -3
)

func (dp DogfightPosition) String() string {
	switch dp {
	case DogfightPositionBehindEnemiesTailOptimal:
		return "BehindEnemyOptimal"
	case DogfightPositionBehindEnemiesTail:
		return "BehindEnemy"
	case DogfightPositionAdventage:
		return "Adventage"
	case DogfightPositionTossup:
		return "Tossup"
	case DogfightPositionDisadvantage:
		return "Disadventage"
	case DogfightPositionEnemyAtMySix:
		return "EnemyAtMySix"
	case DogfightPositionEnemyAtMySixOptimal:
		return "EnemyAtMySixOptimal"
	}
	return "Invalid"
}

// DogfightResult wird nach einem Dogfight zurückgegeben und beinhaltet die Schäden, die an einem Flugzeug verursacht
// wurden.
type DogfightResult struct {
	Round            int
	Position         DogfightPosition
	WeaponUsed       *WeaponSystem
	Hit              bool
	DamageConflicted []DamageType
}

//goland:noinspection ALL
func (dr DogfightResult) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "%dr.ExecuteDogfight Result: ", dr.Round)
	fmt.Fprintf(&sb, "%s", dr.Position)
	if dr.WeaponUsed != nil {
		fmt.Fprintf(&sb, ", HitWith: %s", dr.WeaponUsed.WeaponSystemName)
		if len(dr.DamageConflicted) > 0 {
			fmt.Fprint(&sb, ", DamageConflicted: ")
			for _, dt := range dr.DamageConflicted {
				fmt.Fprintf(&sb, "%s, ", dt.String())
			}
		} else {
			fmt.Fprint(&sb, ", no damage")
		}
	}
	fmt.Fprint(&sb, "\n")
	return sb.String()
}

// DogfightSetup wird initial für einen Kampf benötigt. Aus diesem struct entsteht dann Dogfight.
type DogfightSetup struct {
	TeamBlue AircraftIdList
	TeamRed  AircraftIdList
}

func NewDogfightSetup() DogfightSetup {
	ds := DogfightSetup{}
	ds.TeamBlue = []AircraftId{}
	ds.TeamRed = []AircraftId{}
	return ds
}

func (ds *DogfightSetup) AddRed(id AircraftId) {
	ds.TeamRed = append(ds.TeamRed, id)
}

func (ds *DogfightSetup) AddBlue(id AircraftId) {
	ds.TeamBlue = append(ds.TeamBlue, id)
}

// DogfightGroup wird aus dem struct Dogfight erstellt. Je mehr Flugzeuge in der Dogfight Warteliste sind, desto mehr
// DogfightGroup-Objekte werden erzeugt.
type DogfightGroup struct {
	BlueFighter AircraftId
	BlueFighterLastPosition DogfightPosition
	BlueSupport *AircraftId // optional
	RedFighter  AircraftId
	RedFighterLastPosition DogfightPosition
	RedSupport  *AircraftId // optional
}

func NewDogfightGroup(blue AircraftId, red AircraftId) DogfightGroup {
	dg := DogfightGroup{
		BlueFighter: blue,
		RedFighter:  red,
	}
	dg.BlueFighterLastPosition = DogfightPositionTossup
	dg.RedFighterLastPosition = DogfightPositionTossup
	return dg
}

type DogfightGroupList []DogfightGroup

func (dgl DogfightGroupList) BlueHasFreeSupportSlot() bool {
	for _, group := range dgl {
		if group.HasBlueSupport() == false {
			return true
		}
	}
	return false
}

func (dgl DogfightGroupList) RedHasFreeSupportSlot() bool {
	for _, group := range dgl {
		if group.HasRedSupport() == false {
			return true
		}
	}
	return false
}

func (dgl *DogfightGroupList) AssignBlueSupport(id AircraftId) bool {
	for i, group := range *dgl {
		// Erste gefundene Gruppe mit freiem Support belegen und dann raus hier.
		if group.HasBlueSupport() == false {
			(*dgl)[i].BlueSupport = &id
			return true
		}
	}
	return false
}

func (dgl *DogfightGroupList) AssignRedSupport(id AircraftId) bool {
	for _, group := range *dgl {
		// Erste gefundene Gruppe mit freiem Support belegen und dann raus hier.
		if group.HasRedSupport() == false {
			group.RedSupport = &id
			return true
		}
	}
	return false
}

func (dg DogfightGroup) HasBlueSupport() bool {
	return dg.BlueSupport != nil
}

func (dg DogfightGroup) HasRedSupport() bool {
	return dg.RedSupport != nil
}

// Dogfight wird aus einem DogfightSetup initialisiert. Während des Kampfes werden so viele DogfightGroup erstellt, wie
// es möglich ist.
type Dogfight struct {
	Groups          DogfightGroupList
	TeamBlueWaiting AircraftIdList
	TeamRedWaiting  AircraftIdList
}

func (d *Dogfight) Execute() {

}

// DistributeAircraftsToGroups verteilt wartende Aircrafts auf die Gruppen. Liefert true, wenn min. ein
// Aircraft verteilt werden konnte, sonst false.
func (d *Dogfight) DistributeAircraftsToGroups() bool {
	distributionHappened := false
	// 2er-Gruppen erzeugen
	for len(d.TeamBlueWaiting) > 0 && len(d.TeamRedWaiting) > 0 {
		b := d.TeamBlueWaiting.PullFirst()
		r := d.TeamRedWaiting.PullFirst()
		d.Groups = append(d.Groups, NewDogfightGroup(b, r))
		distributionHappened = true
	}

	if len(d.Groups) > 0 {
		// Vorhandene Gruppen mit restlichen Aircrafts auffüllen.
		for len(d.TeamBlueWaiting) > 0 && d.Groups.BlueHasFreeSupportSlot() {
			b := d.TeamBlueWaiting.PullFirst()
			if d.Groups.AssignBlueSupport(b) == false {
				panic("error while finding free slot for support")
			}
			distributionHappened = true
		}
		for len(d.TeamRedWaiting) > 0 && d.Groups.RedHasFreeSupportSlot() {
			b := d.TeamRedWaiting.PullFirst()
			if d.Groups.AssignRedSupport(b) == false {
				panic("error while finding free slot for support")
			}
			distributionHappened = true
		}
	}
	return distributionHappened
}

func SimulateDogfightPosition(rating1 Rating, lastPosition1 DogfightPosition,
	rating2 Rating, lastPosition2 DogfightPosition) DogfightPosition {
	dfr1 := randomizer.Roll1D10() + int(rating1)
	dfr2 := randomizer.Roll1D10() + int(rating2)
	dfdelta := dfr1 - dfr2

	if lastPosition1 == DogfightPositionAdventage {
		dfdelta = dfdelta + 3
	}

	if lastPosition2 == DogfightPositionAdventage {
		dfdelta = dfdelta - 3
	}

	if dfdelta > 0 {
		if dfdelta >= 7 {
			return DogfightPositionBehindEnemiesTailOptimal
		} else if dfdelta >= 4 {
			return DogfightPositionBehindEnemiesTail
		} else if dfdelta >= 2 {
			return DogfightPositionAdventage
		}
	} else {
		if -dfdelta >= 7 {
			return DogfightPositionEnemyAtMySixOptimal
		} else if -dfdelta >= 4 {
			return DogfightPositionEnemyAtMySix
		} else if -dfdelta >= 2 {
			return DogfightPositionDisadvantage
		}
	}
	return DogfightPositionTossup
}

func (ds DogfightSetup) CreateDogfight() Dogfight {
	d := Dogfight{}
	d.Groups = []DogfightGroup{}
	d.TeamRedWaiting = make(AircraftIdList, len(ds.TeamRed))
	copy(d.TeamRedWaiting, ds.TeamRed)
	d.TeamBlueWaiting = make(AircraftIdList, len(ds.TeamBlue))
	copy(d.TeamBlueWaiting, ds.TeamBlue)
	return d
}

// ExecuteDogfight Eine Runde im Luftkampf. Etwa 10 Sekunden dauer.
func ExecuteDogfight(
	acid1 AircraftId, ldp1 DogfightPosition,
	acid2 AircraftId, ldp2 DogfightPosition) (DogfightResult, DogfightResult) {
	var dfr1 DogfightResult
	var dfr2 DogfightResult
	ac1 := Globals.AllAircrafts[acid1]
	ac2 := Globals.AllAircrafts[acid2]
	ap1 := ac1.GetParameters()
	ap2 := ac2.GetParameters()

	// In FloatPosition setzen
	// Flugzeuge mit grösseren Dogfighting-Rating haben höhere Chance.
	// 1) Kampf um die FloatPosition => Endet in einer FloatPosition
	dfa1Pos := SimulateDogfightPosition(ap1.Dogfighting, ldp1, ap2.Dogfighting, ldp2)
	dfr1.Position = dfa1Pos
	dfr2.Position = -dfa1Pos

	// SRMs (Short-Range-Missles) gegeneinander einsetzen
	// 2) Abschuss der SRM
	// Falls keine SRM => Einsatz der Gun
	if dfa1Pos >= DogfightPositionBehindEnemiesTail {
		bestws, exist := ac1.GetBestDogfightingWeapon()
		dfr1.WeaponUsed = &bestws
		if exist {
			ac1.DepleteWeapon(bestws)
			if bestws.Hit(acid2, dfa1Pos) {
				dfr1.Hit = true
				dt := ac2.DoDamageWith(bestws)
				dfr1.DamageConflicted = append(dfr1.DamageConflicted, dt)
			}
		}
	} else if -dfa1Pos >= DogfightPositionBehindEnemiesTail {
		bestws, exist := ac2.GetBestDogfightingWeapon()
		dfr2.WeaponUsed = &bestws
		if exist {
			ac2.DepleteWeapon(bestws)
			if bestws.Hit(acid1, -dfa1Pos) {
				dfr2.Hit = true
				dt := ac1.DoDamageWith(bestws)
				dfr2.DamageConflicted = append(dfr2.DamageConflicted, dt)
			}
		}
	}
	return dfr1, dfr2
}

func Sim10Rounds(acid1 AircraftId, acid2 AircraftId) (*[]DogfightResult, *[]DogfightResult) {
	drl1 := make([]DogfightResult, 0)
	drl2 := make([]DogfightResult, 0)

	ldp1 := DogfightPositionTossup
	ldp2 := DogfightPositionTossup

	for i := 0; i < 10; i++ {
		dr1, dr2 := ExecuteDogfight(acid1, ldp1, acid2, ldp2)
		dr1.Round = i
		dr2.Round = i
		ldp1 = dr1.Position
		ldp2 = dr2.Position
		drl1 = append(drl1, dr1)
		drl2 = append(drl2, dr2)
	}
	return &drl1, &drl2
}
