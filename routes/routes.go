package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvents)
	server.PUT("/events/:id", updateEvent) // Update
	server.DELETE("/events/:id", deleteEvent) // Delete
	server.POST("/signup",) // 註冊新使用者
	server.GET("/login",) // 使用者登陸
}