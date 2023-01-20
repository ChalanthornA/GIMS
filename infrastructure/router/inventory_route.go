package router

import (
	"github.com/ChalanthornA/Gold-Inventory-Management-System/controllers"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/infrastructure/database"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/middlewares"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/repositories"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/usecases"
	"github.com/gin-gonic/gin"
)

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
		inventory.PATCH("/deletegolddetail", goldController.SetStatusGoldDetailToDelete)
		inventory.PATCH("/getbackgolddetail/:id", goldController.SetStatusGoldDetailToNormal)
		inventory.PATCH("/setgoldinventorystatus", goldController.SetStatusGoldInventory)
		inventory.POST("/getgolddetailjoininventorybydetail", goldController.GetGoldDetailJoinInventoryByDetail)
		inventory.GET("/getgolddetailbygolddetailid/:id", goldController.GetGoldDetailByGoldDetailID)
		inventory.POST("/delete-gold-inventory-by-id", goldController.DeleteGoldInventoryByIDArray)
		inventory.PATCH("/set-tag-serial-number", goldController.SetTagSerialNumberGoldInventory)
	}
}