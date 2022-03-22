package game

func (p *Pilot) insert() error {
	if Globals.Database == nil {
		return DatabaseNotOpenError{}
	}
	tx, err := Globals.Database.Begin()
	if err != nil {
		return err
	}
	stmt, err3 := tx.Prepare("insert into table_pilots(pilot_name, country_name, gender, flight_rank, age, born, home_air_base, sorties, hits, kills, kia, mia, xp, reflexes, endurance) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err3 != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer stmt.Close()
	res, err2 := stmt.Exec(p.Name, p.CountryName, p.Gender.String(), p.FlightRank.Code, p.Background.Age, p.Background.Born, p.Background.HomeAirBase, p.Sorties, p.Hits, p.Kills, p.Kia, p.Mia, p.XP, p.Reflexes, p.Endurance)
	if err2 != nil {
		return err
	}
	var errInsert error
	p.DatabaseId, errInsert = res.LastInsertId()
	if errInsert != nil {
		return errInsert
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
	stmt, err := tx.Prepare("update table_pilots SET pilot_name = ?, country_name = ?, gender = ?, flight_rank = ?, age = ?, born = ?, home_air_base = ?, sorties = ?, hits = ?, kills = ?, kia = ?, mia = ?, xp = ?, reflexes = ?, endurance = ? WHERE pilot_id == ?")
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer stmt.Close()
	_, err = stmt.Exec(p.Name,
		p.CountryName,
		p.Gender.String(),
		p.FlightRank.Code,
		p.Background.Age,
		p.Background.Born,
		p.Background.HomeAirBase,
		p.Sorties, p.Hits, p.Kills, p.Kia, p.Mia, p.XP, p.Reflexes, p.Endurance, p.DatabaseId)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func LoadPilot(pid int64) (Pilot, error) {
	p := Pilot{}
	if Globals.Database == nil {
		return p, DatabaseNotOpenError{}
	}
	stmt, err := Globals.Database.Prepare("select pilot_name, country_name, gender, flight_rank, age, born, home_air_base, sorties, hits, kills, kia, mia, xp, reflexes, endurance from table_pilots WHERE pilot_id = ?")
	if err != nil {
		return p, err
	}

	var gender string
	var frank int
	err = stmt.QueryRow(pid).Scan(
		&p.Name,
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
	p.DatabaseId = pid
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
    pilot_id      integer not null PRIMARY KEY 
);
insert into table_pilots(pilot_name, country_name, gender, flight_rank, age, born, home_air_base, sorties,
                                hits, kills, kia, mia, xp, reflexes, endurance, pilot_id)
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
       pilot_id
from table_pilots;
create unique index table_pilots_id_uindex
    on table_pilots (pilot_id);`

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

func GetPilotsOfCountry(country string) ([]Pilot, error) {
	pilots := make([]Pilot, 0)
	if Globals.Database == nil {
		return pilots, DatabaseNotOpenError{}
	}
	//goland:noinspection GoUnhandledErrorResult
	rows, errq := Globals.Database.Query("SELECT pilot_name, pilot_id, country_name, gender, flight_rank, age, born, home_air_base, sorties, hits, kills, kia, mia, xp, reflexes, endurance FROM table_pilots WHERE country_name = '" + country + "'")
	if errq != nil {
		return pilots, errq
	}
	for rows.Next() {
		p := Pilot{}
		var gender string
		var frank int
		errs := rows.Scan(&p.Name,
			&p.DatabaseId,
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
		if errs != nil {
			return pilots, errs
		}
		if gender == "Male" {
			p.Gender = GenderMale
		} else {
			p.Gender = GenderFemale
		}
		p.FlightRank = NewRank(p.CountryName, (Code)(frank))
		pilots = append(pilots, p)
	}

	return pilots, nil
}
