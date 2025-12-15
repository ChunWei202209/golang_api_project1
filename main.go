package main

import (
	"net/http"

	"example.com/golang-api-project1/db"
	"example.com/golang-api-project1/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvents)

	server.Run(":8080") // localhost:8080
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "無法找到event"})
	}
	context.JSON(http.StatusOK, events)
}

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