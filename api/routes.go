package api

import (
	"example.com/golang-api-project1/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// 建立一個「路由群組」，方便統一加 middleware 或共用設定
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvents)
	authenticated.PUT("/events/:id", updateEvent) // Update
	authenticated.DELETE("/events/:id", deleteEvent) // Delete
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)
	
	// server.POST("/events", middlewares.Authenticate, createEvents)
	// server.PUT("/events/:id", updateEvent) 
	// server.DELETE("/events/:id", deleteEvent) 

	server.POST("/signup", signup) // 使用者註冊
	server.POST("/login", login) // 使用者登錄
}



