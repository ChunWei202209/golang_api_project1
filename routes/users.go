package routes

import (
	"net/http"

	"example.com/golang-api-project1/models"
	"example.com/golang-api-project1/utils"
	"github.com/gin-gonic/gin"
)

// 創造使用者
func signup(context *gin.Context) {
	var user models.User

	// ShouldBindJSON: 把 request body 的 JSON 轉成 Go 的 struct
	// 如果轉換失敗，就把錯誤存進 err
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "無法處理 user"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "無法創造 user"})
		return
	}

	// 把整個 user 包成 JSON 回給 client
	// 包含自動生成的 ID、前面填好的 UserId，以及 client 原本送的資料
	context.JSON(http.StatusCreated, gin.H{"message": "事件被創造", "user": user})
}

// 登錄
func login(context *gin.Context) {
	var user models.User

	// ShouldBindJSON: 把 request body 的 JSON 轉成 Go 的 struct
	// 如果轉換失敗，就把錯誤存進 err
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "無法處理 user"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		println(err.Error())
		context.JSON(http.StatusUnauthorized, gin.H{"message": "無法認證使用者"})
		return
	}

	// 呼叫 JWT
	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		println("GenerateToken error:", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": "無法認證使用者!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "登陸成功!", "token": token})
}