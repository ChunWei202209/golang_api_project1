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

// 從 events 資料表把每一列撈出來 → 
// 轉成 Event struct → 
// 收集成 slice 回傳
func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	// rows：一個「游標」，一列一列讀
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 宣告一個「裝 Event 的盒子列隊」
	// 因為 Scan 每一列都要有一個暫存容器
	var events []Event

	// 迴圈來讀取每一行資料
	for rows.Next() {
		var event Event // 幫我生一個 全新的 Event 變數
		
		// 資料庫這一列的每一個欄位，
		// 請你直接寫進這個 event 裡對應的欄位。
		// Scan 要的是「位址」
		if err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserId,
		); err != nil {
			return nil, err
		}
		// 把 目前這個 event 的「值」 複製一份，
		// 放進 events 這個 slice 裡。
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}