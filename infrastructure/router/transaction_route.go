package router

import (
	"github.com/ChalanthornA/Gold-Inventory-Management-System/controllers"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/infrastructure/database"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/repositories"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/usecases"
	"github.com/gin-gonic/gin"
)

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
		transaction.GET("/gettransactionbytimeinterval/:range", transactionController.GetTransactionByTimeInterval)
		transaction.GET("/get-transaction-by-date/:date", transactionController.GetTransactionByDate)
		transaction.GET("/get-transaction-from-to", transactionController.GetTransactionFromTo)
		transaction.GET("get-daily-report", transactionController.GetDailyReport)
	}
}