package router

import (
	"github.com/ChalanthornA/Gold-Inventory-Management-System/controllers"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/infrastructure/database"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/middleware"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/repositories"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/usecases"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine){
	userRepository := repositories.NewUserRepository(database.DB)
	userUseCase := usecases.NewUserUseCase(userRepository)
	userController := controllers.NewUserController(userUseCase)

	users := r.Group("/users")
	users.POST("/registeradmin", userController.RegisterAdmin)
	users.POST("/signin", userController.SignIn)

	users.Use(middleware.AuthorizeJWT())
	users.POST("/register", userController.Register)
	users.GET("/testjwt", userController.TestJWT)
}