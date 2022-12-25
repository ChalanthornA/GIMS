package domains

import "github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"

type TransactionUseCase interface {
	NewTransactionTypeBuy(transaction *models.Transaction) error
	NewTransactionTypeSell(transaction *models.Transaction) error
	NewTransactionTypeChange(transaction *models.Transaction) error
	RollBackTransaction(transactionID uint32) error
	GetAllTransactionJoinGold() ([]models.TransactionJoinGold, error)
}

type TransactionRepository interface {
	InsertNewTransaction(transaction *models.Transaction) error
	DeleteTransaction(transactionID uint32) error
	QueryTransactionByTransactionID(transactionID uint32) (*models.Transaction, error)
	QueryAllTransaction() ([]models.TransactionJoinGold, error)
}