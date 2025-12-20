package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// 需要權限才能使用
	server.POST("/events", createEvents)
	server.PUT("/events/:id", updateEvent) // Update
	server.DELETE("/events/:id", deleteEvent) // Delete

	server.POST("/signup", signup) // 使用者註冊
	server.POST("/login", login) // 使用者登錄
}