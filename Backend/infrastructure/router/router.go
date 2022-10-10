package router

import (
	"github.com/ChalanthornA/Gold-Inventory-Management-System/controllers"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/infrastructure/database"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/middleware"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/repositories"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/usecases"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App){
	userRepository := repositories.NewUserRepository(database.DB)
	userUseCase := usecases.NewUserUseCase(userRepository)
	userController := controllers.NewUserController(userUseCase)

	users := app.Group("/users")
	users.Post("/registeradmin", userController.RegisterAdmin)
	users.Post("/signin", userController.SignIn)
	
	users.Use(middleware.AuthorizationRequired())
	users.Get("/testjwt", userController.TestJWT)
}