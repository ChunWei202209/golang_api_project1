package main

import (
	"net/http"

	"example.com/golang-api-project1/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvents)

	server.Run(":8080") // localhost:8080
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) {
	var event models.Event

	// ShouldBindJSON: 把 request body 的 JSON 轉成 Go 的 struct
	// 如果轉換失敗，就把錯誤存進 err
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "無法處理請求"})
		return
	}

	event.ID = 1
	event.UserId = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "事件被創造", "event": event})
}