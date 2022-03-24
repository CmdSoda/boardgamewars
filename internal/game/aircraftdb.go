package game

func CreateAircraftTable() error {
	var stmt string
	stmt = `create table table_aircrafts
(
    aircraft_id      integer PRIMARY KEY AUTOINCREMENT,
    parked_at_airbase_id integer,
    maintained_at_airbase_id integer,
	pos_x integer,
	pos_y integer
);
insert into table_aircrafts(aircraft_id, parked_at_airbase_id, maintained_at_airbase_id, pos_x, pos_y)
select aircraft_id,
       parked_at_airbase_id,
       maintained_at_airbase_id,
       pos_x,
       pos_y
from table_aircrafts;
create unique index table_aircrafts_id_uindex
    on table_aircrafts (aircraft_id);`

	_, err := Globals.Database.Exec(stmt)
	return err
}

func DropAircraftTable() error {
	if Globals.Database == nil {
		return DatabaseNotOpenError{}
	}

	//goland:noinspection GoUnhandledErrorResult
	Globals.Database.Exec("DROP TABLE IF EXISTS table_aircrafts")

	return nil
}
