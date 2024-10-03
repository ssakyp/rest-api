package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		//crashes the app if there is not DB connection
		panic("Could not connect to database.")
	}

	// established connection
	// ta manage the connection
	DB.SetMaxOpenConns(10) // how many db connections we can have open simultaneously

	DB.SetMaxIdleConns(5) // how many connections we keep open if noone is using it

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table.")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table.")
	}

	// 1. create table
	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER
		FOREIGN KEY(event_id) REFERENCE events(id)
		FOREIGN KEY(user_id) REFERENCE users(id)
	)
	`
	// 2. execute it
	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic("Could not create registrations table.")
	}
}
