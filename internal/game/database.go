package game

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type DatabaseNotOpenError struct {
}

func (dberr DatabaseNotOpenError) Error() string {
	return "Database not open"
}

func OpenDatabase() error {
	db, err := sql.Open("sqlite3", "../../data/bgw.sqlite")
	if err != nil {
		return err
	}
	Globals.Database = db
	return nil
}

func CloseDatabase() {
	//goland:noinspection GoUnhandledErrorResult
	Globals.Database.Close()
	Globals.Database = nil
}