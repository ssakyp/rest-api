package models

import (
	"time"

	"github.com/pelletier/go-toml/query"
	"github.com/ssakyp/rest-api/db"
)

// this package is for storing event data in DB, fetching data and so on

// data / shape that makes up the event
type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

var events []Event = []Event{} // empty slice of Events

func (e Event) Save() error {
	// new query we target the table, in () different fields then values
	// ? special syntax gives us a sql injection save way
	query := `INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	// we can pass as many as values
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err
	// later: add it to a database
	// events = append(events, e)	// no longer locally
}

// to fetch the data
func GetAllEvents() []Event {
	query := "SELECT * FROM events"
	db.DB.Query(query)
	return events
}
