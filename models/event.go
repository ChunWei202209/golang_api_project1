package models

// Models：
// 跟資料庫溝通，存資料、拿資料、更新資料。

// 做三件事：
// 1️⃣ 定義資料長什麼樣（struct）。
// 2️⃣ 定義資料怎麼進 DB（Save / Update）。
// 3️⃣ 定義資料怎麼出 DB（GetAll / GetByID），需要使用地址寫入。

// 執行步驟：
// 1. 將 SQL 指令存進變數
// 2. 根據用途選擇合適的方法呼叫 DB：
//    - Exec → 不需要回傳資料，只送資料給 DB（INSERT/UPDATE/DELETE）
//    - Query → 可能回傳多筆資料，需用 Rows.Next() 讀取
//    - QueryRow → 只回傳一筆資料
// 3. 若回傳的是 *Rows，使用完一定要 defer rows.Close()；QueryRow/Exec 不需要 Close
// 4. 若需要把 DB 回傳的資料寫進 Go 變數（struct 或其他），使用 Scan 並傳地址；否則不用 Scan
// 	1) 宣告一個變數（例如 struct 或普通變數），準備裝資料
// 	2) 使用 `Scan(&var1, &var2, ...)`，把每個欄位寫進對應變數
//    - **注意**：Scan 需要位址（`&`），因為它要把資料寫進變數
//    - 方向：DB → Go 記憶體
// 5. 回傳必要資料或錯誤

// 方法解釋：
// - DB.Prepare(...) 適合同一條 SQL 需要重複執行（多筆資料或多次 request）;
// - DB.Exec(...)  適合單次執行 SQL，用於建表、修改資料；
// - DB.QueryRow(...)  用於查詢一筆資料；
// - DB.Query(...)  用於查詢多筆資料。不需要呼叫 Close()，因為 QueryRow 內部已經幫你管理了連線與資源;

// 地址說明：
// Scan 是「寫 Go 變數」的工具
// Exec 是「只把資料送進 DB」的工具
// 有把「DB 回傳的值」寫進 Go struct → 一定要 Scan，而且要傳地址
// 沒有把值寫回 Go 變數 → 用 Exec，不用 Scan

import (
	"time"

	"example.com/golang-api-project1/internal/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserID      int64     `json:"userId"`
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
		e.UserID)
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
		var event Event // 幫我生一個 全新的 Event 變數 來裝回傳值
		
		// 資料庫這一列的每一個欄位，
		// 直接寫進這個 event 裡對應的 struct 欄位。
		// Scan 要的是「位址」
		if err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserID,
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

	var event Event // 幫我生一個 全新的 Event 變數 來裝回傳值

	// 資料庫這一列的每一個欄位，
	// 直接寫進這個 event 裡對應的 struct 欄位。
	// Scan 要的是「位址」
	
	err := row.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserID,
		) 
		
		if err != nil {
			return nil, err
		}

		return &event, nil
}

// 更新 - 改「內容」 → struct 是內容集合
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

// 刪除 - 刪「存在本身」 → ID 就夠了
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

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	stmt.Exec(e.ID, userId)

	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id =? AND user_id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	stmt.Exec(e.ID, userId)

	return err
}