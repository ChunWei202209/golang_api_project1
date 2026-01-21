package utils

// JWT 可以把使用者的身分資訊（例如 userId）
// 變成一段被簽名過的字串，讓前端帶著它來證明「我是誰」。

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

// 1️⃣ 建立一張「還沒蓋章的身分證」
// 回傳 token 還有可能失敗的 error
// - 在使用者成功登入 / 註冊後呼叫
// - 把 email、userId 等資訊放進 JWT
// - 使用 secret key 簽名
// - 回傳一個可以給前端保存的 token 字串
func GenerateToken(email string, userId int64) (string, error) {
	fmt.Println(userId)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"exp": time.Now().Add(time.Hour *2).Unix(), // 現在時間 + 2 小時
	})

	// 2️⃣ 用 secret key 蓋章，變成真的 JWT
	return token.SignedString([]byte(secretKey))
}

// VerifyToken 用來驗證傳入的 JWT token 是否有效
// - 簽名正確（沒有被竄改）
// - 尚未過期
func VerifyToken(token string) (int64, error) {
	// jwt.Parse 會解析傳入的 token 並驗證簽名
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			// 如果不是 HMAC，回傳錯誤
			return nil, errors.New("預期之外的登錄方法")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("無法處理 token")
	}
	
	// 檢查 token 是否有效（簽名正確且未過期）
	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("無效的 token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("無效的 token claims")
	}

	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return userId, nil
}