package usecases

import "github.com/ChalanthornA/Gold-Inventory-Management-System/domains"

type transactionUsecase struct {
	transactionRepo domains.TransactionRepository
	goldRepo domains.GoldRepository
}

func NewTransactionUsecase(tr domains.TransactionRepository, gr domains.GoldRepository) domains.TransactionUseCase {
	return &transactionUsecase{transactionRepo: tr, goldRepo: gr}
}