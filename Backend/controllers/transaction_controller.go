package controllers

import (
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains"
)

type transactionController struct {
	transactionUseCase domains.TransactionUseCase
}

func NewTransactionController(tu domains.TransactionUseCase) *transactionController{
	return &transactionController{transactionUseCase: tu}
}