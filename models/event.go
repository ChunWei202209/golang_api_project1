package models

// Models：
// 跟資料庫溝通，存資料、拿資料、更新資料。

// 做三件事：
// 1️⃣ 定義資料長什麼樣（struct）。
// 2️⃣ 定義資料怎麼進 DB（Save / Update）。
// 3️⃣ 定義資料怎麼出 DB（GetAll / GetByID）。

// 一句話結論（給新手用的版本）
// 有 SELECT → 一定要 Scan
// 沒有 SELECT → 用 Exec，不 Scan
// Scan = 把 DB 的欄位值寫進 Go 變數

import (
	"time"

	"example.com/golang-api-project1/db"
)

// 資料的藍圖
type Event struct {
	ID 					int64  `json:"id"`
	Name 				string `binding:"required"` // 「前端送 JSON 時，這個欄位一定要有」
	Description string `binding:"required"`
	Location 		string `binding:"required"`
	DateTime 		time.Time `binding:"required"`
	UserId 			int
}

// 把 Event 存進資料庫，使用指標來修改 ID
func (e *Event) Save() error {

	// 使用 ?，避免 SQL Injection
	query := `
		INSERT INTO events(name, description, location, dateTime, user_id) 
		VALUES (?, ?, ?, ?, ?)
	`
	// Prepare：先把 SQL 準備好
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	// 「這個 function 結束時，幫我把 stmt 關掉」避免資源外洩。
	defer stmt.Close()

	// 真的執行 INSERT
	result, err := stmt.Exec(
		e.Name, 
		e.Description, 
		e.Location, 
		e.DateTime, 
		e.UserId)
	if err != nil {
		return err
	}

	// 拿回自動產生的 ID
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

// 取得所有資料
func GetAllEvents() ([]Event, error) {

	// 從 events 資料表把每一列撈出來 → 
	// 轉成 Event struct → 
	// 收集成 slice 回傳
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

// 取回 ID
func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE ID = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserId,
		) 
		
		if err != nil {
			return nil, err
		}

		return &event, nil
}

// 更新 - 👉 改「內容」 → struct 是內容集合
func (event Event) Update() error {
	query := `
		UPDATE events
		SET name = ?, description = ?, location = ?, dateTime = ?
		WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		event.Name, 
		event.Description,
		event.Location, 
		event.DateTime, 
		event.ID,
	)
	return err
}

// 刪除 - 👉 刪「存在本身」 → ID 就夠了
func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err
}
