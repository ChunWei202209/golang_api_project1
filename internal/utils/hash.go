package utils

// 用途：
// - 將「使用者輸入的原始密碼」轉成不可逆的雜湊字串
// - 寫入資料庫的永遠是 hash，不是原始密碼

import "golang.org/x/crypto/bcrypt"

// 加密
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// 用途：
// - 驗證「登入時輸入的密碼」是否和 DB 中的 hash 相符
// 驗證密碼是否正確
func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}