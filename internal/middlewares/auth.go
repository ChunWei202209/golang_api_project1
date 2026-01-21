package middlewares

// 這個檔案負責「身分驗證」相關的 middleware。

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"example.com/golang-api-project1/internal/utils"
	"example.com/golang-api-project1/internal/logger"
)

func Authenticate(context *gin.Context) {
	// 從 HTTP Header 讀取 Authorization token
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "沒有權限"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		logger.Log.Warn("認證失敗", logger.ErrorField(err))
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "沒有權限"})
		return
	}

	// 將已驗證的 userId 存入 gin.Context，
	// 讓後續的 handler 可以直接取得目前登入的使用者
	context.Set("userId", userId)

	// middleware 工作完成，放行請求，
	// 繼續執行下一個 middleware 或 handler
	context.Next()
}