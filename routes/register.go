package routes

import (
	"net/http"
	"strconv"

	"example.com/golang-api-project1/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) 
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "無法取得 ID"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "無法取得 event"})
	  return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "無法註冊使用者到活動"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "註冊成功!"})

}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) 

	var event models.Event
	event.ID = eventId
	
	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "無法取消註冊使用者到活動"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "取消註冊成功!"})
}