package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
	"strconv"
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
	Round                   int
	Position                DogfightPosition
	WeaponUsed              *WeaponSystem
	Hit                     bool
	DamageConflictedToEnemy []DamageType
}

//goland:noinspection ALL
func (dr DogfightResult) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "%dr.ExecuteDogfight Result: ", dr.Round)
	fmt.Fprintf(&sb, "%s", dr.Position)
	if dr.WeaponUsed != nil {
		fmt.Fprintf(&sb, ", HitWith: %s", dr.WeaponUsed.WeaponSystemName)
		if len(dr.DamageConflictedToEnemy) > 0 {
			fmt.Fprint(&sb, ", DamageConflictedToEnemy: ")
			for _, dt := range dr.DamageConflictedToEnemy {
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

var currentDogfightGroupShortId ShortId = 0

// DogfightGroup wird aus dem struct Dogfight erstellt. Je mehr Flugzeuge in der Dogfight Warteliste sind, desto mehr
// DogfightGroup-Objekte werden erzeugt.
type DogfightGroup struct {
	ShortId                 ShortId
	BlueFighterId           AircraftId
	BlueFighterLastPosition DogfightPosition
	BlueSupportId           *AircraftId // optional
	RedFighterId            AircraftId
	RedFighterLastPosition  DogfightPosition
	RedSupportId            *AircraftId // optional
}

func (dg *DogfightGroup) Simulate() (DogfightResult, DogfightResult) {
	var dfr1 DogfightResult
	var dfr2 DogfightResult

	ac1 := Globals.AllAircrafts[dg.BlueFighterId]
	ap1 := ac1.GetParameters() // AircraftParameters

	ac2 := Globals.AllAircrafts[dg.RedFighterId]
	ap2 := ac2.GetParameters() // AircraftParameters

	// In FloatPosition setzen
	// Flugzeuge mit grösseren Dogfighting-Rating haben höhere Chance.
	// 1) Kampf um die FloatPosition => Endet in einer FloatPosition
	dfa1Pos := SimulateDogfightPosition(ap1.Dogfighting, dg.BlueFighterLastPosition, ap2.Dogfighting, dg.RedFighterLastPosition)
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
			if bestws.Hit(dg.RedFighterId, dfa1Pos) {
				dfr1.Hit = true
				dt := ac2.DoDamageWith(bestws)
				dfr1.DamageConflictedToEnemy = append(dfr1.DamageConflictedToEnemy, dt)
			}
		}
	} else if -dfa1Pos >= DogfightPositionBehindEnemiesTail {
		bestws, exist := ac2.GetBestDogfightingWeapon()
		dfr2.WeaponUsed = &bestws
		if exist {
			ac2.DepleteWeapon(bestws)
			if bestws.Hit(dg.BlueFighterId, -dfa1Pos) {
				dfr2.Hit = true
				dt := ac1.DoDamageWith(bestws)
				dfr2.DamageConflictedToEnemy = append(dfr2.DamageConflictedToEnemy, dt)
			}
		}
	}
	return dfr1, dfr2
}

// HasAircraftsOnBothSides liefert true, solange Aircrafts auf beiden Seiten der WarPartys vorhanden sind.
func (dg DogfightGroup) HasAircraftsOnBothSides() bool {
	b := Globals.AllAircrafts[dg.BlueFighterId]
	bAvailable := b.Destroyed == false
	bsAvailable := dg.BlueSupportId != nil
	r := Globals.AllAircrafts[dg.RedFighterId]
	rAvailable := r.Destroyed == false
	rsAvailable := dg.RedSupportId != nil
	return bAvailable || bsAvailable && rAvailable || rsAvailable
}

func NewDogfightGroup(blue AircraftId, red AircraftId) DogfightGroup {
	dg := DogfightGroup{
		BlueFighterId: blue,
		RedFighterId:  red,
	}
	dg.BlueFighterLastPosition = DogfightPositionTossup
	dg.RedFighterLastPosition = DogfightPositionTossup
	dg.ShortId = currentDogfightGroupShortId
	currentDogfightGroupShortId = currentDogfightGroupShortId + 1
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
			(*dgl)[i].BlueSupportId = &id
			Log.Infof("distribute: blue AC%d assigned to group DG%d",
				Globals.AllAircrafts[id].ShortId, group.ShortId)
			return true
		}
	}
	return false
}

func (dgl *DogfightGroupList) AssignRedSupport(id AircraftId) bool {
	for i, group := range *dgl {
		// Erste gefundene Gruppe mit freiem Support belegen und dann raus hier.
		if group.HasRedSupport() == false {
			(*dgl)[i].RedSupportId = &id
			Log.Infof("distribute: red AC%d assigned to group DG%d",
				Globals.AllAircrafts[id].ShortId, group.ShortId)
			return true
		}
	}
	return false
}

func (dg DogfightGroup) HasBlueSupport() bool {
	return dg.BlueSupportId != nil
}

func (dg DogfightGroup) HasRedSupport() bool {
	return dg.RedSupportId != nil
}

func (dg DogfightGroup) String() string {
	bid := strconv.Itoa(int(Globals.AllAircrafts[dg.BlueFighterId].ShortId))
	bsid := "<empty>"
	if dg.BlueSupportId != nil {
		bsid = strconv.Itoa(int(Globals.AllAircrafts[*dg.BlueSupportId].ShortId))
	}
	rid := strconv.Itoa(int(Globals.AllAircrafts[dg.RedFighterId].ShortId))
	rsid := "<empty>"
	if dg.RedSupportId != nil {
		rsid = strconv.Itoa(int(Globals.AllAircrafts[*dg.RedSupportId].ShortId))
	}
	return fmt.Sprintf("DogfightGroup(DG%d): Blue=AC%s (alive=%t) BlueSupport=AC%s, Red=AC%s (alive=%t) RedSupport=AC%s",
		dg.ShortId,
		bid, !Globals.AllAircrafts[dg.BlueFighterId].Destroyed, bsid,
		rid, !Globals.AllAircrafts[dg.RedFighterId].Destroyed, rsid)
}

// Dogfight wird aus einem DogfightSetup initialisiert. Während des Kampfes werden so viele DogfightGroup erstellt, wie
// es möglich ist.
type Dogfight struct {
	Groups          DogfightGroupList
	TeamBlueWaiting AircraftIdList
	TeamRedWaiting  AircraftIdList
}

func (d Dogfight) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "Groups: len=%d ", len(d.Groups))
	if len(d.Groups) == 0 {
		fmt.Fprint(&b, "<empty>\n")
		return b.String()
	} else {
		fmt.Fprintf(&b, "\n")
		for _, group := range d.Groups {
			fmt.Fprintf(&b, "%s\n", group.String())
		}
	}
	return b.String()
}

func (d *Dogfight) Simulate() {
	Log.Infof("starting a dogfight simulation with %d groups", len(d.Groups))
	for i, _ := range d.Groups {
		blueResult, redResult := d.Groups[i].Simulate()
		d.Groups[i].BlueFighterLastPosition = blueResult.Position
		d.Groups[i].RedFighterLastPosition = redResult.Position
	}
}

// DistributeAircraftsToGroups verteilt wartende Aircrafts auf die Gruppen. Liefert true, wenn min. ein
// Aircraft verteilt werden konnte, sonst false.
func (d *Dogfight) DistributeAircraftsToGroups() bool {
	groupsAdded := 0
	aircraftsCount := len(d.TeamBlueWaiting) + len(d.TeamRedWaiting)
	sendToWaitCount := 0
	distributionHappened := false

	restart := true

	for restart {
		restart = false
		// Vorhandene Gruppen aufräumen
		for i, _ := range d.Groups {
			// Wenn BlueFighter zerstört wurde, muss er durch den Support ersetzt werden.
			if Globals.AllAircrafts[d.Groups[i].BlueFighterId].Destroyed && d.Groups[i].BlueSupportId != nil {
				d.Groups[i].BlueFighterId = *d.Groups[i].BlueSupportId
				d.Groups[i].BlueSupportId = nil
				Log.Infof("distribute: blue support AC%d is now lead in the group",
					Globals.AllAircrafts[d.Groups[i].BlueFighterId].ShortId)
			}
			// Das gleiche gilt für Red
			if Globals.AllAircrafts[d.Groups[i].RedFighterId].Destroyed && d.Groups[i].RedSupportId != nil {
				d.Groups[i].RedFighterId = *d.Groups[i].RedSupportId
				d.Groups[i].RedSupportId = nil
				Log.Infof("distribute: red support AC%d is now lead in the group",
					Globals.AllAircrafts[d.Groups[i].RedFighterId].ShortId)
			}

			// Sollte Blue komplett ausgelöscht werden sein, wird Red wieder in die Warteliste getragen und
			// ist bereit für andere Dogfights.
			if Globals.AllAircrafts[d.Groups[i].BlueFighterId].Destroyed && d.Groups[i].BlueSupportId == nil {
				if Globals.AllAircrafts[d.Groups[i].RedFighterId].Destroyed == false {
					d.TeamRedWaiting = append(d.TeamRedWaiting, d.Groups[i].RedFighterId)
					Log.Infof("distribute: red AC%d send to waiting queue",
						Globals.AllAircrafts[d.Groups[i].RedFighterId].ShortId)
					sendToWaitCount = sendToWaitCount + 1
					// Gruppe i löschen
					Log.Infof("distribute: removing DG%d", d.Groups[i].ShortId)
					d.Groups = append(d.Groups[:i], d.Groups[i+1:]...)
					// Wir haben ein Element gelöscht und müssen die for-Schleife von vorn starten, sonst
					// gibt es einen Index-Fehler
					restart = true
					break
				}
			}
			// Das gleiche auch für Red
			if Globals.AllAircrafts[d.Groups[i].RedFighterId].Destroyed && d.Groups[i].RedSupportId == nil {
				if Globals.AllAircrafts[d.Groups[i].BlueFighterId].Destroyed == false {
					d.TeamBlueWaiting = append(d.TeamBlueWaiting, d.Groups[i].BlueFighterId)
					Log.Infof("distribute: blue AC%d send to waiting queue",
						Globals.AllAircrafts[d.Groups[i].BlueFighterId].ShortId)
					sendToWaitCount = sendToWaitCount + 1
					// Gruppe i löschen
					Log.Infof("distribute: removing DG%d", d.Groups[i].ShortId)
					d.Groups = append(d.Groups[:i], d.Groups[i+1:]...)
					// Wir haben ein Element gelöscht und müssen die for-Schleife von vorn starten, sonst
					// gibt es einen Index-Fehler
					restart = true
					break
				}
			}
		}
	}

	// 2er-Gruppen erzeugen
	for len(d.TeamBlueWaiting) > 0 && len(d.TeamRedWaiting) > 0 {
		b := d.TeamBlueWaiting.PullFirst()
		r := d.TeamRedWaiting.PullFirst()
		d.Groups = append(d.Groups, NewDogfightGroup(b, r))
		groupsAdded = groupsAdded + 1
		distributionHappened = true
		Log.Infof("distribute: blue AC%d and red AC%d forming a new group",
			Globals.AllAircrafts[b].ShortId,
			Globals.AllAircrafts[r].ShortId)
	}

	Log.Infof("distribute: distributed %d aircrafts to %d new dogfight groups, sent %d AC to wait queue",
		aircraftsCount, groupsAdded, sendToWaitCount)

	if len(d.Groups) > 0 {
		// Vorhandene Gruppen mit restlichen Aircrafts auffüllen.
		for len(d.TeamBlueWaiting) > 0 && d.Groups.BlueHasFreeSupportSlot() {
			b := d.TeamBlueWaiting.PullFirst()
			if d.Groups.AssignBlueSupport(b) == false {
				panic("distribute: error while finding free slot for support")
			}
			distributionHappened = true
		}
		for len(d.TeamRedWaiting) > 0 && d.Groups.RedHasFreeSupportSlot() {
			b := d.TeamRedWaiting.PullFirst()
			if d.Groups.AssignRedSupport(b) == false {
				panic("distribute: error while finding free slot for support")
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
		if dfdelta >= 5 {
			return DogfightPositionBehindEnemiesTailOptimal
		} else if dfdelta >= 3 {
			return DogfightPositionBehindEnemiesTail
		} else if dfdelta >= 1 {
			return DogfightPositionAdventage
		}
	} else {
		if -dfdelta >= 5 {
			return DogfightPositionEnemyAtMySixOptimal
		} else if -dfdelta >= 3 {
			return DogfightPositionEnemyAtMySix
		} else if -dfdelta >= 1 {
			return DogfightPositionDisadvantage
		}
	}
	return DogfightPositionTossup
}

func NewDogfight(ds DogfightSetup) Dogfight {
	d := Dogfight{}
	d.Groups = []DogfightGroup{}
	d.TeamRedWaiting = make(AircraftIdList, len(ds.TeamRed))
	copy(d.TeamRedWaiting, ds.TeamRed)
	d.TeamBlueWaiting = make(AircraftIdList, len(ds.TeamBlue))
	copy(d.TeamBlueWaiting, ds.TeamBlue)
	return d
}
