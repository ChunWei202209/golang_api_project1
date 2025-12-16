package main

import (
	"example.com/golang-api-project1/db"
	"example.com/golang-api-project1/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB() // 建立資料庫連線
	server := gin.Default() // 建立一個 Gin Engine

	routes.RegisterRoutes(server) // 把 URL 跟 handler function 綁在一起

	server.Run(":8080") // localhost:8080
}

