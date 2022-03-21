package game

import (
	"fmt"
	"github.com/CmdSoda/boardgamewars/internal/randomizer"
	"github.com/google/uuid"
)

type Gender int

const (
	GenderMale   Gender = 0
	GenderFemale Gender = 1
)

func (g Gender) String() string {
	if g == GenderMale {
		return "Male"
	} else {
		return "Female"
	}
}

type PilotId uuid.UUID

type PilotsMap map[PilotId]*Pilot

type Pilot struct {
	PilotId
	Name        string
	CountryName // Gehört diesem Land an
	Gender
	Background PilotBackground
	FlightRank
	PilotStats
}

type PilotStats struct {
	Sorties   int
	Hits      int
	Kills     int
	Kia       bool
	Mia       bool
	XP        int
	Reflexes  int // Reflexe im ExecuteDogfight: -2, -1, 0, +1, +2
	Endurance int // Wieviel Stress kann der Pilot verkraften.
}

func (p Pilot) String() string {
	return fmt.Sprintf("%s(%s) (%s) [%s] (%d-%d-%d) %dyo (Born: %s HomeBase: %s)",
		p.Name,
		p.Gender,
		p.FlightRank,
		p.Background.CountryName,
		p.PilotStats.Sorties,
		p.PilotStats.Hits,
		p.PilotStats.Kills,
		p.Background.Age,
		p.Background.Born,
		p.Background.HomeAirBase)
}

func (p Pilot) Short() string {
	return p.Name
}

// RollAge
// https://www.operationmilitarykids.org/air-force-age-limits/
func RollAge(ofc Code) int {
	switch ofc {
	case OF1:
		return randomizer.Roll(22, 24)
	case OF2:
		return randomizer.Roll(25, 28)
	case OF3:
		return randomizer.Roll(27, 31)
	case OF4:
		return randomizer.Roll(29, 35)
	case OF5:
		return randomizer.Roll(32, 39)
	}
	return 0
}

// NewPilot erzeugt einen neuen Piloten und speichert ihn in die Datenbank.
func NewPilot(country CountryName, ofc Code) *Pilot {
	var g Gender

	wp, exist := Globals.CountryDataMap[country]
	if exist == false {
		panic("country does not exist")
	}

	gr := randomizer.Roll1D10()
	if gr >= 2 {
		g = GenderMale
	} else {
		g = GenderFemale
	}

	ng := Generator{}
	var cd *CountryDataItem
	cd, exist = Globals.CountryDataMap[wp.CountryName]
	if !exist {
		Log.Panicf("data for %s not found", wp.CountryName)
		return nil
	}
	ng.AddNameSet(&cd.NameSet)

	np := Pilot{
		Name:        ng.CreateFullName(g == GenderMale, wp.CountryName),
		CountryName: country,
		PilotId:     PilotId(uuid.New()),
		Gender:      g,
		Background: PilotBackground{
			CountryName: wp.CountryName,
			Born:        ng.CreateCityName(wp.CountryName),
			Age:         RollAge(ofc),
			HomeAirBase: ng.CreateAirForceBaseName(wp.CountryName),
		},
		FlightRank: NewRank(wp.CountryName, ofc),
	}

	Log.Infof("new pilot created: %s", np.Short())
	if errdb := np.insert(); errdb != nil {
		Log.Panicf("Error while saving pilot: %s", errdb)
	}
	return &np
}

func NewPilots(count int, country CountryName, ofc Code) []PilotId {
	var pilots []PilotId
	for i := 0; i < count; i++ {
		np := NewPilot(country, ofc)
		pilots = append(pilots, np.PilotId)
	}
	return pilots
}

func (p *Pilot) insert() error {
	if Globals.Database == nil {
		return DatabaseNotOpenError{}
	}
	tx, err := Globals.Database.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("insert into table_pilots(pilot_name, pilot_uuid, country_name, gender, flight_rank, age, born, home_air_base, sorties, hits, kills, kia, mia, xp, reflexes, endurance) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer stmt.Close()
	uid := (uuid.UUID)(p.PilotId)
	_, err = stmt.Exec(p.Name, uid.String(), p.CountryName, p.Gender.String(), p.FlightRank.Code, p.Background.Age, p.Background.Born, p.Background.HomeAirBase, p.Sorties, p.Hits, p.Kills, p.Kia, p.Mia, p.XP, p.Reflexes, p.Endurance)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (p *Pilot) Update() error {
	if Globals.Database == nil {
		return DatabaseNotOpenError{}
	}
	tx, err := Globals.Database.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("update table_pilots SET pilot_name = ?, country_name = ?, gender = ?, flight_rank = ?, age = ?, born = ?, home_air_base = ?, sorties = ?, hits = ?, kills = ?, kia = ?, mia = ?, xp = ?, reflexes = ?, endurance = ? WHERE pilot_uuid == ?")
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer stmt.Close()
	uid := (uuid.UUID)(p.PilotId)
	_, err = stmt.Exec(p.Name,
		p.CountryName,
		p.Gender.String(),
		p.FlightRank.Code,
		p.Background.Age,
		p.Background.Born,
		p.Background.HomeAirBase,
		p.Sorties, p.Hits, p.Kills, p.Kia, p.Mia, p.XP, p.Reflexes, p.Endurance, uid.String())
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func LoadPilot(pid PilotId) (Pilot, error) {
	p := Pilot{}
	if Globals.Database == nil {
		return p, DatabaseNotOpenError{}
	}
	stmt, err := Globals.Database.Prepare("select pilot_name, pilot_uuid, country_name, gender, flight_rank, age, born, home_air_base, sorties, hits, kills, kia, mia, xp, reflexes, endurance from table_pilots WHERE pilot_uuid = ?")
	if err != nil {
		return p, err
	}

	var gender string
	var pilotUuid string
	var frank int
	err = stmt.QueryRow((uuid.UUID)(pid).String()).Scan(
		&p.Name,
		&pilotUuid,
		&p.CountryName,
		&gender,
		&frank,
		&p.Background.Age,
		&p.Background.Born,
		&p.Background.HomeAirBase,
		&p.Sorties,
		&p.Hits,
		&p.Kills,
		&p.Kia,
		&p.Mia,
		&p.XP,
		&p.Reflexes,
		&p.Endurance)
	if err != nil {
		return p, err
	}
	p.PilotId = pid
	if gender == "Male" {
		p.Gender = GenderMale
	} else {
		p.Gender = GenderFemale
	}
	p.FlightRank = NewRank(p.CountryName, (Code)(frank))

	return p, nil
}

func CreatePilotTable() error {
	var stmt string
	stmt = `create table table_pilots
(
    pilot_name    string  not null,
    country_name  string  not null,
    gender        string  not null,
    flight_rank   integer,
    age           integer,
    born          string,
    home_air_base string,
    sorties       integer,
    hits          integer,
    kills         integer,
    kia           boolean,
    mia           boolean,
    xp            integer not null,
    reflexes      integer,
    endurance     integer,
    pilot_uuid    string not null PRIMARY KEY 
);
insert into table_pilots(pilot_name, country_name, gender, flight_rank, age, born, home_air_base, sorties,
                                hits, kills, kia, mia, xp, reflexes, endurance, pilot_uuid)
select pilot_name,
       country_name,
       gender,
       flight_rank,
       age,
       born,
       home_air_base,
       sorties,
       hits,
       kills,
       kia,
       mia,
       xp,
       reflexes,
       endurance,
       pilot_uuid
from table_pilots;
create unique index table_pilots_uuid_uindex
    on table_pilots (pilot_uuid);`

	_, err := Globals.Database.Exec(stmt)
	return err
}

func DropPilotTable() error {
	if Globals.Database == nil {
		return DatabaseNotOpenError{}
	}

	//goland:noinspection GoUnhandledErrorResult
	Globals.Database.Exec("DROP TABLE IF EXISTS table_pilots")

	return nil
}

func RemoveAllPilots() error {
	if Globals.Database == nil {
		return DatabaseNotOpenError{}
	}
	//goland:noinspection GoUnhandledErrorResult,SqlWithoutWhere
	Globals.Database.Exec("DELETE FROM table_pilots")

	return nil
}

func NumberOfPilots() (int, error) {
	if Globals.Database == nil {
		return 0, DatabaseNotOpenError{}
	}
	//goland:noinspection GoUnhandledErrorResult,SqlWithoutWhere
	query, err := Globals.Database.Prepare("SELECT COUNT(*) FROM table_pilots")
	if err != nil {
		return 0, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer query.Close()
	var count int
	err = query.QueryRow().Scan(&count)

	return count, err
}
