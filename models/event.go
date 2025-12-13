package models

import "time"

type Event struct {
	ID 					int
	Name 				string `binding:"required"`
	Description string `binding:"required"`
	Location 		string `binding:"required"`
	DateTime 		time.Time `binding:"required"`
	UserId 			int
}

var events = []Event{}

func (e Event) Save() {
	// add it to a database

	// append「加到 slice 後面」的工具。
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}