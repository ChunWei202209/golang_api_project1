package api

import (
	"net/http"
	"strconv"

	"example.com/golang-api-project1/models"
	"example.com/golang-api-project1/internal/logger"
	"github.com/gin-gonic/gin"
)

// @Summary 報名活動
// @Security Bearer
// @Param id path int true "活動ID"
// @Success 201 {object} map[string]interface{}
// @Router /events/{id}/register [post]
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
		logger.Log.Error("報名活動失敗", logger.ErrorField(err))
		context.JSON(http.StatusInternalServerError, gin.H{"message": "無法註冊使用者到活動"})
		return
	}
	logger.Log.Info("用戶報名活動成功", logger.Int64Field("eventId", eventId), logger.Int64Field("userId", userId))
	context.JSON(http.StatusCreated, gin.H{"message": "註冊成功!"})

}

// @Summary 取消報名
// @Security Bearer
// @Param id path int true "活動ID"
// @Success 200 {object} map[string]interface{}
// @Router /events/{id}/register [delete]
func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) 
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "無法取得 ID"})
		return
	}

	var event models.Event
	event.ID = eventId
	
	err = event.CancelRegistration(userId)

	if err != nil {
		logger.Log.Error("取消報名失敗", logger.ErrorField(err))
		context.JSON(http.StatusInternalServerError, gin.H{"message": "無法取消註冊使用者到活動"})
		return
	}
	logger.Log.Info("取消報名活動成功", logger.Int64Field("eventId", eventId), logger.Int64Field("userId", userId))
	context.JSON(http.StatusOK, gin.H{"message": "取消報名活動成功!"})
}

