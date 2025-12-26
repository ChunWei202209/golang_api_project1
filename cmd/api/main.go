package main

import (
	"example.com/golang-api-project1/api"
	"example.com/golang-api-project1/internal/db"
	"example.com/golang-api-project1/internal/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "example.com/golang-api-project1/docs" // swagger docs
)

// @title     Golang API Project
// @version   1.0
// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	logger.InitLogger()
	defer logger.Sync()

	db.InitDB()
	server := gin.Default()
	api.RegisterRoutes(server)

	// Swagger 路由
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := server.Run(":8080"); err != nil {
		logger.Log.Fatal("伺服器啟動失敗", logger.ErrorField(err))
	}
}

