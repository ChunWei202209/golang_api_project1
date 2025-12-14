package models

import (
	"time"

	"example.com/golang-api-project1/db"
)

type Event struct {
	ID 					int64
	Name 				string `binding:"required"`
	Description string `binding:"required"`
	Location 		string `binding:"required"`
	DateTime 		time.Time `binding:"required"`
	UserId 			int
}

var events = []Event{}

func (e *Event) Save() error {
	// 避免 SQL Injection
	query := `
		INSERT INTO events(name, description, location, dateTime, user_id) 
		VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	// 真的執行 INSERT
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}

	// 自動產生 id
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() []Event {
	return events
}