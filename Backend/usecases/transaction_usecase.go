package usecases

import (
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
)

type transactionUsecase struct {
	transactionRepo domains.TransactionRepository
	goldRepo domains.GoldRepository
}

func NewTransactionUsecase(tr domains.TransactionRepository, gr domains.GoldRepository) domains.TransactionUseCase {
	return &transactionUsecase{transactionRepo: tr, goldRepo: gr}
}

func (tu *transactionUsecase) NewTransactionTypeBuy(transaction *models.Transaction) error {
	transaction.SetTimeNow()
	return tu.transactionRepo.InsertNewTransaction(transaction)
}

func (tu *transactionUsecase) NewTransactionTypeSell(transaction *models.Transaction) error {
	goldInventory, err := tu.goldRepo.CheckGoldInventoryByGoldInventoryID(transaction.GoldInventoryID)
	if err != nil {
		return err
	}
	goldDetail, err := tu.goldRepo.QueryGoldDetailByGoldDetailID(goldInventory.GoldDetailID)
	if err != nil {
		return err
	}
	transaction.GoldDetailID = goldInventory.GoldDetailID
	transaction.Weight = goldDetail.Weight
	transaction.SetTimeNow()
	if err := tu.goldRepo.UpdateGoldInventoryStatus(transaction.GoldInventoryID, "sold"); err != nil {
		return err
	}
	return tu.transactionRepo.InsertNewTransaction(transaction)
}

func (tu *transactionUsecase) NewTransactionTypeChange(transaction *models.Transaction) error {
	goldInventory, err := tu.goldRepo.CheckGoldInventoryByGoldInventoryID(transaction.GoldInventoryID)
	if err != nil {
		return err
	}
	if _, err := tu.goldRepo.QueryGoldDetailByGoldDetailID(goldInventory.GoldDetailID); err != nil {
		return err
	}
	transaction.GoldDetailID = goldInventory.GoldDetailID
	transaction.Price = transaction.SellPrice - transaction.BuyPrice
	transaction.SetTimeNow()
	if err := tu.goldRepo.UpdateGoldInventoryStatus(transaction.GoldInventoryID, "sold"); err != nil {
		return err
	}
	return tu.transactionRepo.InsertNewTransaction(transaction)
}