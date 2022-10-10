package main

import (
	"github.com/ChalanthornA/Gold-Inventory-Management-System/infrastructure/database"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/infrastructure/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.DB = database.NewDb()
	defer database.DB.Close()

	router.SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":8080")
}