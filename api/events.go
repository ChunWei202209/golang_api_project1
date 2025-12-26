package api

// Controller：

// 做四件事：
// 1️⃣ 收到客戶端來的請求，Gin 找到對應 handler。
// 2️⃣ 呼叫 model 層去拿資料。
// 3️⃣ 檢查是否出錯。
// 4️⃣ 回傳資料給 client。

// context *gin.Context 是 Gin 框架給每一個 HTTP 請求的一個 指標核心物件，
// 裡面裝了整個請求需要的東西。

import (
	"net/http"
	"strconv"

	"example.com/golang-api-project1/models"
	"example.com/golang-api-project1/internal/logger"
	"github.com/gin-gonic/gin"
)

// @Summary 取得所有活動
// @Success 200 {array} models.Event
// @Router /events [get]
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "無法找到event"})
		return
	}
	logger.Log.Info("成功取得所有活動", logger.Int64Field("count", int64(len(events))))
	// 如果沒有錯誤，就把從 DB 拿到的 events 轉成 JSON 回給 client
	context.JSON(http.StatusOK, events)
}

// @Summary 取得單一活動
// @Param id path int true "活動ID"
// @Success 200 {object} models.Event
// @Router /events/{id} [get]
func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // 轉成 int64
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "無法取得 ID"})
		return
	}

	event ,err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "無法取得 event"})
		return
	}

	logger.Log.Info("成功取得活動", logger.Int64Field("eventId", eventId))
	context.JSON(http.StatusOK, event)
}

// @Summary 創造活動
// @Security Bearer
// @Param event body models.Event true "活動資訊"
// @Success 201 {object} map[string]interface{}
// @Router /events [post]
func createEvents(context *gin.Context) {

	// 宣告一個「空的 Event 容器」，型別是 models.Event，也就是 struct，
	// 預設欄位都是 空 或 0
	var event models.Event

	// ShouldBindJSON: 把 request body 的 JSON 轉成 Go 的 struct
	// 如果轉換失敗，就把錯誤存進 err
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "無法處理 event"})
		return
	}
	userId := context.GetInt64("userId")
	event.UserID = userId

	// 傳給 model 層存資料
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "無法創造 event"})
		return
	}

	logger.Log.Info("成功創造活動", logger.Int64Field("eventId", event.ID), logger.Int64Field("userId", userId), logger.StringField("eventName", event.Name))
	// 把整個 event 包成 JSON 回給 client
	// 包含自動生成的 ID、前面填好的 UserId，以及 client 原本送的資料
	context.JSON(http.StatusCreated, gin.H{"message": "事件被創造", "event": event})
}

// @Summary 更新活動
// @Security Bearer
// @Param id path int true "活動ID"
// @Param event body models.Event true "活動資訊"
// @Success 200 {object} map[string]interface{}
// @Router /events/{id} [put]
func updateEvent(context *gin.Context) {
	// 先檢查 ID 存不存在
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) 
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "無法取得 ID"})
		return
	}

	// 取得 ID
	// 因為只有使用者才能編輯活動，所以要驗證 ID 是不是相同
	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "無法取得 event"})
		return
	}

	// 若 ID 不符合
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "沒有權限去更新 event"})
		return
	}

	// 宣告一個「空的 Event 容器」，型別是 models.Event，也就是 struct，
	// 預設欄位都是 空 或 0
	var updatedEvent models.Event

	// 把 request body 的 JSON 轉成 Go 的 struct
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
	})
		return
	}

	updatedEvent.ID = eventId

	// 傳給 model 層更新資料
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "無法更新 event"})
		return
	}
	logger.Log.Info("成功更新活動", logger.Int64Field("eventId", eventId), logger.Int64Field("userId", userId))
	context.JSON(http.StatusOK, gin.H{"message": "成功更新 event"})

}

// @Summary 刪除活動
// @Security Bearer
// @Param id path int true "活動ID"
// @Success 200 {object} map[string]interface{}
// @Router /events/{id} [delete]
func deleteEvent(context *gin.Context) {
	// 先檢查 ID 存不存在
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) 
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "無法取得 ID"})
		return
	}

	// 取得 ID
	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "無法更新 event"})
		return
	}

	// 若 ID 不符合
	// 因為只有使用者才能刪除活動，所以要驗證 ID 是不是相同
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "沒有權限去刪除 event"})
		return
	}

	err = event.Delete()

	if err != nil {
		logger.Log.Error("刪除活動失敗", logger.ErrorField(err))
		context.JSON(http.StatusInternalServerError, gin.H{"message": "無法刪除 event"})
		return
	}
	logger.Log.Info("成功刪除活動", logger.Int64Field("eventId", eventId), logger.Int64Field("userId", userId))
	context.JSON(http.StatusOK, gin.H{"message": "成功刪除 event"})
}

