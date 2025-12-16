package main

import (
	"example.com/golang-api-project1/db"
	"example.com/golang-api-project1/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}

