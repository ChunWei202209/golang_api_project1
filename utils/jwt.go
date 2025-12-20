package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

// 1️⃣ 建立一張「還沒蓋章的身分證」
// 回傳 token 還有可能失敗的 error
func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"exp": time.Now().Add(time.Hour *2).Unix(), // 現在時間 + 2 小時
	})

	// 2️⃣ 用 secret key 蓋章，變成真的 JWT
	return token.SignedString([]byte(secretKey))
}