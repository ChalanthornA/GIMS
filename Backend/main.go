package main

import (
	"github.com/ChalanthornA/Gold-Inventory-Management-System/infrastructure/database"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/infrastructure/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.DB = database.NewDb()
	defer database.DB.Close()

	r.Use(cors.Default())
	router.SetupRoutes(r)

	r.Run()
}