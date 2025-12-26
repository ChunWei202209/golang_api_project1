package api

import (
	"net/http"

	"example.com/golang-api-project1/models"
	"example.com/golang-api-project1/internal/utils"
	"example.com/golang-api-project1/internal/logger"
	"github.com/gin-gonic/gin"
)

// @Summary 使用者註冊
// @Param user body models.User true "使用者資訊"
// @Success 201 {object} map[string]interface{}
// @Router /signup [post]
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
		logger.Log.Warn("註冊失敗", logger.StringField("email", user.Email), logger.ErrorField(err))
		context.JSON(http.StatusBadRequest, gin.H{"message": "無法創造 user"})
		return
	}

	logger.Log.Info("用戶註冊成功", logger.StringField("email", user.Email), logger.Int64Field("userId", user.ID))
	// 把整個 user 包成 JSON 回給 client
	// 包含自動生成的 ID、前面填好的 UserId，以及 client 原本送的資料
	context.JSON(http.StatusCreated, gin.H{"message": "事件被創造", "user": user})
}

// @Summary 使用者登錄
// @Param user body models.User true "登錄資訊"
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func login(context *gin.Context) {
	var user models.User

	// ShouldBindJSON: 把 request body 的 JSON 轉成 Go 的 struct
	// 如果轉換失敗，就把錯誤存進 err
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "無法處理 user"})
		return
	}

	// 驗整使用者
	err = user.ValidateCredentials()

	if err != nil {
		logger.Log.Warn("登入失敗", logger.StringField("email", user.Email))
		context.JSON(http.StatusUnauthorized, gin.H{"message": "無法認證使用者"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		logger.Log.Error("無法生成 token", logger.ErrorField(err))
		context.JSON(http.StatusInternalServerError, gin.H{"message": "無法認證使用者!"})
		return
	}
	logger.Log.Info("用戶登入成功", logger.StringField("email", user.Email), logger.Int64Field("userId", user.ID))
	context.JSON(http.StatusOK, gin.H{"message": "用戶登入成功", "token": token})
}

