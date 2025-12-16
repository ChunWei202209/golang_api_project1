package routes

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
	}
	context.JSON(http.StatusOK, events)
}

// 取得 ID
func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) 
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
func createEvents(context *gin.Context) {
	var event models.Event

	// ShouldBindJSON: 把 request body 的 JSON 轉成 Go 的 struct
	// 如果轉換失敗，就把錯誤存進 err
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "無法處理 event"})
		return
	}

	event.ID = 1
	event.UserId = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "無法創造 event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "事件被創造", "event": event})
}