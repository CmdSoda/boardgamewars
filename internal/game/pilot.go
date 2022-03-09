package game

import (
	"database/sql"
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
	DatabaseId int64
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

	Globals.CountryDataMap[country].Pilots = append(Globals.CountryDataMap[country].Pilots, np.PilotId)
	Globals.AllPilots[np.PilotId] = &np
	Log.Infof("new pilot created: %s", np.Short())
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

func (p *Pilot) Save() error {
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
	var res sql.Result
	uid := (uuid.UUID)(p.PilotId)
	res, err = stmt.Exec(p.Name, uid.String(), p.CountryName, p.Gender.String(), p.FlightRank.Code, p.Background.Age, p.Background.Born, p.Background.HomeAirBase, p.Sorties, p.Hits, p.Kills, p.Kia, p.Mia, p.XP, p.Reflexes, p.Endurance)
	if err != nil {
		return err
	}
	p.DatabaseId, err = res.LastInsertId()
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (p *Pilot) Load() error {
	uid := (uuid.UUID)(p.PilotId)

	stmt, err := Globals.Database.Prepare("select pilot_name, pilot_uuid, country_name, gender, flight_rank, age, born, home_air_base, sorties, hits, kills, kia, mia, xp, reflexes, endurance from table_pilots WHERE pilot_uuid = ?")
	if err != nil {
		return err
	}

	var gender string
	var pilotUuid string
	var frank int
	err = stmt.QueryRow(uid.String()).Scan(
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
		return err
	}
	pid, err2 := uuid.Parse(pilotUuid)
	if err2 != nil {
		return err2
	}
	p.PilotId = (PilotId)(pid)
	if gender == "Male" {
		p.Gender = GenderMale
	} else {
		p.Gender = GenderFemale
	}
	p.FlightRank = NewRank(p.CountryName, (Code)(frank))

	return nil
}

func CreatePilotTable() error {
	var stmt string
	stmt = `create table table_pilots
(
    id            integer not null
        constraint table_pilots_pk
            primary key autoincrement,
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
    pilot_uuid    string
);
insert into table_pilots(id, pilot_name, country_name, gender, flight_rank, age, born, home_air_base, sorties,
                                hits, kills, kia, mia, xp, reflexes, endurance, pilot_uuid)
select id,
       pilot_name,
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
create unique index table_pilots_id_uindex
    on table_pilots (id);`

	_, err := Globals.Database.Exec(stmt)
	return err
}

func DropPilotTable() {
	//goland:noinspection GoUnhandledErrorResult
	Globals.Database.Exec("DROP TABLE IF EXISTS table_pilots")
}
