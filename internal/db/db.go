package db

// DB Package：
// 負責初始化資料庫、建立資料表，並提供全域 DB 連線給程式使用。

// 做三件事：
// 1️⃣ 管理資料庫檔案（路徑、資料夾、連線）
// 2️⃣ 初始化資料表（CREATE TABLE IF NOT EXISTS）
// 3️⃣ 提供全域變數 DB，讓其他 package 可以操作資料庫

// 使用規範：
// - 所有資料操作請透過 DB 物件進行
// - Exec 用於建表、修改資料；Query 用於查詢

import (
	"database/sql"
	"os"

	"example.com/golang-api-project1/internal/logger"
	_ "github.com/glebarez/sqlite"
)

var DB *sql.DB

func InitDB() {
	// 資料庫檔案統一放在 data/api.db
	dbPath := "data/api.db"

	// 如果 data 資料夾不存在就建立
	os.MkdirAll("data", 0755)

	// 連接資料庫
	var err error
	DB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		logger.Log.Fatal("無法連到 DB 資料庫", logger.ErrorField(err))
		panic("無法連到 DB 資料庫。" + err.Error())
	}
	// 管理連線池
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

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
	// Exec 是 執行一段 SQL 指令，但不會回傳查詢結果
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		logger.Log.Fatal("無法創造 users table", logger.ErrorField(err))
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
	// Exec 是 執行一段 SQL 指令，但不會回傳查詢結果
	_, err = DB.Exec(createEventsTable)

	if err != nil {
		logger.Log.Fatal("無法創造 events table", logger.ErrorField(err))
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
	// Exec 是 執行一段 SQL 指令，但不會回傳查詢結果
	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		logger.Log.Fatal("無法創造 registrations table", logger.ErrorField(err))
		panic("無法創造 registratinos table." + err.Error())
	}
}
