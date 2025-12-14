package db

import (
  "database/sql"
  _ "github.com/glebarez/sqlite"
)

var DB *sql.DB
 
func InitDB() {
    var err error
    DB, err = sql.Open("sqlite", "api.db")
 
    if err != nil {
        panic("無法連到 DB 資料庫。")
    }
		
    DB.SetMaxOpenConns(10) // 連線池同時的最高數量
    DB.SetMaxIdleConns(5) // 持續開放的連線數目

		createTables()
}

// 這是一個初始化用的函式，通常只在程式啟動時呼叫一次
func createTables() {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER
		)
	`
	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("無法創造 events table." + err.Error())
	}
}