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

	auth := r.Group("/auth")
	{
		auth.POST("/registeradmin", userController.RegisterAdmin)
		auth.POST("/signin", userController.SignIn)

		auth.Use(middleware.AuthorizeJWT())
		auth.POST("/register", userController.Register)
		auth.GET("/profile", userController.TestJWT)
	}
	
	goldRepository := repositories.NewGoldRepository(database.DB)
	goldUseCase := usecases.NewGoldUseCase(goldRepository)
	goldController := controllers.NewGoldController(goldUseCase)

	inventory := r.Group("inventory")
	{
		inventory.POST("/newgold", goldController.NewGold)
	}
}