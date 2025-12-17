package models

import "example.com/golang-api-project1/db"

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

type User struct {
	ID 				int64
	Email 		string `binding:"required"`
	Password 	string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	
	// 「這個 function 結束時，幫我把 stmt 關掉」避免資源外洩。
	defer stmt.Close()

	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}