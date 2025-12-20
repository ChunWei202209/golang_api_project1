package utils

import (
	"errors"
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

// VerifyToken 用來驗證傳入的 JWT token 是否有效
func VerifyToken(token string) error {
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
		return errors.New("無法處理 token")
	}
	
	// 檢查 token 是否有效（簽名正確且未過期）
	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return errors.New("無效的 token")
	}

	// claims, ok := parsedToken.Claims.(jtw.MapClaims)

	// if !ok {
	// 	return errors.New("無效的 token claims")
	// }

	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)

	return nil
}