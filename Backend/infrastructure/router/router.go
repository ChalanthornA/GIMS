package router

import (
	"github.com/ChalanthornA/Gold-Inventory-Management-System/controllers"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/infrastructure/database"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/middlewares"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/repositories"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/usecases"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.Use(middlewares.CORSMiddleware())
	userRepository := repositories.NewUserRepository(database.DB)
	userUseCase := usecases.NewUserUseCase(userRepository)
	userController := controllers.NewUserController(userUseCase)
	auth := r.Group("/auth")
	{
		auth.POST("/registeradmin", userController.RegisterAdmin)
		auth.POST("/signin", userController.SignIn)

		auth.Use(middlewares.AuthorizeJWT())
		auth.POST("/register", userController.Register)
		auth.GET("/profile", userController.TestJWT)
		auth.PUT("/renameusername", userController.RenameUsername)
		auth.PUT("/updatepassword", userController.UpdatePassword)
		auth.GET("/queryalluser", userController.QueryAllUser)
		auth.DELETE("/deleteuser/:username", userController.DeleteUser)
	}

	r.Use(middlewares.AuthorizeJWT())

	goldRepository := repositories.NewGoldRepository(database.DB)
	goldUseCase := usecases.NewGoldUseCase(goldRepository)
	goldController := controllers.NewGoldController(goldUseCase)
	inventory := r.Group("/inventory")
	{
		inventory.POST("/newgold", goldController.NewGold)
		inventory.GET("/findgolddetailbycode/:code", goldController.FindGoldDetailByCode)
		inventory.POST("/findgolddetailbydetail", goldController.FindGoldDetailByDetail)
	}
}
