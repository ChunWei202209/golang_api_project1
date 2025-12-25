package main

import (
	"example.com/golang-api-project1/api"
	"example.com/golang-api-project1/internal/db"
	"example.com/golang-api-project1/internal/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.InitLogger(true)
	defer logger.Sync()

	db.InitDB()
	server := gin.Default()
	api.RegisterRoutes(server)

	if err := server.Run(":8080"); err != nil {
		logger.Log.Fatal("伺服器啟動失敗", logger.ErrorField(err))
	}
}

