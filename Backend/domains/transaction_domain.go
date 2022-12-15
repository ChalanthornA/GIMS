package domains

import "github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"

type TransactionUseCase interface {
	NewTransactionTypeBuy(transaction *models.Transaction) error
	NewTransactionTypeSell(transaction *models.Transaction) error
}

type TransactionRepository interface {
	InsertNewTransaction(transaction *models.Transaction) error
}