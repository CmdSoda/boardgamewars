package game

func CreateAirbaseTable() error {
	var stmt string
	stmt = `create table table_airbases
(
    airbase_id      integer PRIMARY KEY AUTOINCREMENT, 
    airbase_name    string  not null,
	belongs_to		string not null,
	accept_allies   boolean,
	max_parkings_slots integer,
	max_maintenance_slots integer,
	pos_x integer,
	pos_y integer
);
insert into table_airbases(airbase_id, airbase_name, belongs_to, accept_allies, max_parkings_slots, max_maintenance_slots, pos_x, pos_y)
select airbase_id,
       airbase_name,
       belongs_to,
       accept_allies,
       max_parkings_slots,
       max_maintenance_slots,
       pos_x,
       pos_y
from table_airbases;
create unique index table_airbases_id_uindex
    on table_airbases (airbase_id);`

	_, err := Globals.Database.Exec(stmt)
	return err
}

func DropAirbaseTable() error {
	if Globals.Database == nil {
		return DatabaseNotOpenError{}
	}

	//goland:noinspection GoUnhandledErrorResult
	Globals.Database.Exec("DROP TABLE IF EXISTS table_airbases")

	return nil
}
