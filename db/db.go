package db

import (
  "database/sql"
	"os"
  _ "github.com/glebarez/sqlite"
)

var DB *sql.DB
 
func InitDB() {
  // 2. 優先讀取環境變數 DB_PATH，如果沒設定，就用原本的 "api.db"
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "api.db"
	}

	var err error
	// 3. 使用 dbPath 代替寫死的檔案名稱
	DB, err = sql.Open("sqlite", dbPath)

	if err != nil {
		panic("無法連到 DB 資料庫。" + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

// 這是一個初始化用的函式，通常只在程式啟動時呼叫一次
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
		panic("無法創造 users table." + err.Error())
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
		panic("無法創造 events table." + err.Error())
	}

	createRegistrationsTable := `
		CREATE TABLE IF NOT EXISTS registrations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_id INTEGER,
			user_id INTEGER,
			FOREIGN KEY(event_id) REFERENCES events(id),
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`
	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic("無法創造 registratinos table." + err.Error())
	}
}