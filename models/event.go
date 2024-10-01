package models

import "time"

// this package is for storing event data in DB, fetching data and so on

// data / shape that makes up the event
type Event struct {
	ID          int       `binding: "required"`
	Name        string    `binding: "required"`
	Description string    `binding: "required"`
	Location    string    `binding: "required"`
	DateTime    time.Time `binding: "required"`
	UserID      int
}

var events []Event = []Event{} // empty slice of Events

func (e Event) Save() {
	// later: add it to a database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
