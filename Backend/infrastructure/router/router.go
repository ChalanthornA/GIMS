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
	setUpAuthRoute(r)

	r.Use(middlewares.AuthorizeJWT())
	setUpInventoryRoute(r)
	setUpTransactionRoute(r)
}

func setUpAuthRoute(r *gin.Engine) {
	userRepository := repositories.NewUserRepository(database.DB, database.GormDB)
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
}

func setUpInventoryRoute(r *gin.Engine) {
	goldRepository := repositories.NewGoldRepository(database.DB, database.GormDB)
	goldUseCase := usecases.NewGoldUseCase(goldRepository)
	goldController := controllers.NewGoldController(goldUseCase)
	inventory := r.Group("/inventory")
	{
		inventory.GET("/findgolddetailbycode/:code", goldController.FindGoldDetailByCode)
		inventory.POST("/findgolddetailbydetail", goldController.FindGoldDetailByDetail)

		inventory.Use(middlewares.AuthorizeAdminOrOwner())
		inventory.POST("/newgold", goldController.NewGold)
		inventory.POST("/addgold", goldController.AddGold)
		inventory.GET("/getallgolddetail", goldController.GetAllGoldDetail)
		inventory.PUT("/editgolddetail", goldController.EditGoldDetail)
		inventory.GET("/getalldetailjoininventory", goldController.GetAllGoldDetailJoinInventory)
		inventory.PATCH("/deletegolddetail/:id", goldController.SetStatusGoldDetailToDelete)
		inventory.PATCH("/getbackgolddetail/:id", goldController.SetStatusGoldDetailToNormal)
		inventory.PATCH("/setgoldinventorystatus", goldController.SetStatusGoldInventory)
	}
}

func setUpTransactionRoute(r *gin.Engine) {
	goldRepository := repositories.NewGoldRepository(database.DB, database.GormDB)
	transactionRepository := repositories.NewTransactionRepository(database.DB, database.GormDB)
	transactionUsecase := usecases.NewTransactionUsecase(transactionRepository, goldRepository)
	transactionController := controllers.NewTransactionController(transactionUsecase)
	transaction := r.Group("/transaction")
	{
		transaction.POST("/newbuytransaction", transactionController.NewTransactionBuy)
		transaction.POST("/newselltransaction", transactionController.NewTransactionSell)
		transaction.POST("/newchangetransaction", transactionController.NewTransactionChange)
		transaction.POST("/rollbacktransaction", transactionController.RollBackTransaction)
		transaction.GET("/getalltransactionjoingold", transactionController.GetAllTransactionJoinGold)
		transaction.GET("/gettransactionbytransactiontype", transactionController.GetAllTransactionByTransactionType)
	}
}