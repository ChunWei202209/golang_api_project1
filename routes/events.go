package routes

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
	"github.com/gin-gonic/gin"
)

// 取得所有活動
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "無法找到event"})
		return
	}
	// 如果沒有錯誤，就把從 DB 拿到的 events 轉成 JSON 回給 client
	context.JSON(http.StatusOK, events)
}

// 依照 URL param 的 id 取得單一 event
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

	context.JSON(http.StatusOK, event)
}

// 創造活動
// 需要權限才能使用，所以需要加入驗證 token
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

	// 把整個 event 包成 JSON 回給 client
	// 包含自動生成的 ID、前面填好的 UserId，以及 client 原本送的資料
	context.JSON(http.StatusCreated, gin.H{"message": "事件被創造", "event": event})
}

// 更新
func updateEvent(context *gin.Context) {
	// 先檢查 ID 存不存在
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) 
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "無法取得 ID"})
		return
	}

	// 取得 ID
	_, err = models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "無法取得 event"})
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
	context.JSON(http.StatusOK, gin.H{"message": "成功更新 event"})

}

func deleteEvent(context *gin.Context) {
	// 先檢查 ID 存不存在
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) 
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "無法取得 ID"})
		return
	}

	// 取得 ID
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "無法更新 event"})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "無法刪除 event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "成功刪除 event"})
}