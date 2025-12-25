package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"example.com/golang-api-project1/utils"
	"example.com/golang-api-project1/logger"
)

func Authenticate(context *gin.Context) {
	// 驗證
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

	context.Set("userId", userId)
	context.Next()
}