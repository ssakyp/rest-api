package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	DB, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		//crashes the app if there is not DB connection
		panic("Could not connect to database.")
	}

	// established connection
	// ta manage the connection
	DB.SetMaxOpenConns(10) // how many db connections we can have open simultaneously

	DB.SetMaxIdleConns(5) // how many connections we keep open if noone is using it
}
