package main

import (
	"github.com/ChalanthornA/Gold-Inventory-Management-System/pkg/common/db"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/pkg/users"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db := db.InitDb()

	users.RegisterUserRoutes(app, db)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":8080")
}