package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

type handler struct{
	DB *pgxpool.Pool
}

func RegisterUserRoutes(app *fiber.App, db *pgxpool.Pool){
	h := &handler{
		DB: db,
	}

	users := app.Group("/users")
	users.Post("/registeradmin", h.RegisterAdmin)
}