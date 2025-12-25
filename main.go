package main

import (
	"example.com/golang-api-project1/db"
	"example.com/golang-api-project1/logger"
	"example.com/golang-api-project1/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.InitLogger(true)
	defer logger.Sync()

	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	if err := server.Run(":8080"); err != nil {
		logger.Log.Fatal("伺服器啟動失敗", logger.ErrorField(err))
	}
}

