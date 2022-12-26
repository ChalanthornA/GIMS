package repositories

import (
	"context"
	"fmt"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/infrastructure/database"
	"github.com/jackc/pgx/v4/pgxpool"
	"gorm.io/gorm"
)

type transactionRepository struct {
	db     *pgxpool.Pool
	gormDb *gorm.DB
	ctx    context.Context
}

func NewTransactionRepository(db *pgxpool.Pool, gormDb *gorm.DB) domains.TransactionRepository {
	return &transactionRepository{db: db, ctx: context.Background(), gormDb: gormDb}
}

func (tr *transactionRepository) InsertNewTransaction(transaction *models.Transaction) error {
	insertTransactionSql := `
		INSERT INTO transactions (
			transaction_id,
			transaction_type,
			date,
			gold_price,
			weight,
			price,
			gold_detail_id,
			gold_inventory_id,
			username,
			buy_price,
			sell_price,
			note
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
		);
	`
	_, err := tr.db.Exec(tr.ctx, insertTransactionSql, database.GenerateUUID(), transaction.TransactionType, transaction.Date, transaction.GoldPrice, transaction.Weight, transaction.Price, transaction.GoldDetailID, transaction.GoldInventoryID, transaction.Username, transaction.BuyPrice, transaction.SellPrice, transaction.Note)
	return err
}

func (tr *transactionRepository) QueryTransactionByTransactionID(transactionID uint32) (*models.Transaction, error) {
	transaction := new(models.Transaction)
	queryTransactionByTransactionIDSql := `SELECT * FROM transactions WHERE transaction_id = $1;`
	row := tr.db.QueryRow(tr.ctx, queryTransactionByTransactionIDSql, transactionID)
	if err := row.Scan(&transaction.TransactionID, &transaction.TransactionType, &transaction.Date, &transaction.GoldPrice, &transaction.Weight, &transaction.Price, &transaction.GoldDetailID, &transaction.GoldInventoryID, &transaction.Username, &transaction.BuyPrice, &transaction.SellPrice, &transaction.Note); err != nil {
		return transaction, err
	}
	if transaction.TransactionID == 0 {
		return transaction, fmt.Errorf("can't find transaction")
	}
	return transaction, nil
}

func (tr *transactionRepository) DeleteTransaction(transactionID uint32) error {
	deleteTransactionSql := `DELETE FROM transactions WHERE transaction_id = $1;`
	_, err := tr.db.Exec(tr.ctx, deleteTransactionSql, transactionID)
	return err
}

func (tr *transactionRepository) QueryAllTransaction() ([]models.TransactionJoinGold, error) {
	var transactionJoinGolds []models.TransactionJoinGold
	queryAllTransactionSql := `SELECT * FROM transactions;`
	rows, err := tr.db.Query(tr.ctx, queryAllTransactionSql)
	if err != nil {
		return transactionJoinGolds, err
	}
	for rows.Next() {
		var transactionJoinGold models.TransactionJoinGold
		err = rows.Scan(&transactionJoinGold.Transaction.TransactionID, &transactionJoinGold.Transaction.TransactionType, &transactionJoinGold.Transaction.Date, &transactionJoinGold.Transaction.GoldPrice, &transactionJoinGold.Transaction.Weight, &transactionJoinGold.Transaction.Price, &transactionJoinGold.Transaction.GoldDetailID, &transactionJoinGold.Transaction.GoldInventoryID, &transactionJoinGold.Transaction.Username, &transactionJoinGold.Transaction.BuyPrice, &transactionJoinGold.Transaction.SellPrice, &transactionJoinGold.Transaction.Note)
		if err != nil {
			return transactionJoinGolds, err
		}
		transactionJoinGolds = append(transactionJoinGolds, transactionJoinGold)
	}
	return transactionJoinGolds, nil
}

func (tr *transactionRepository) QueryTransactionByTransactionType(transactionType string) ([]models.TransactionJoinGold, error) {
	var transactions []models.TransactionJoinGold
	queryTransactionByTransactionTypeSql := `SELECT * FROM transactions WHERE transaction_type = ?;`
	rows, err := tr.gormDb.Raw(queryTransactionByTransactionTypeSql, transactionType).Rows()
	if err != nil {
		return transactions, err
	}
	for rows.Next() {
		var transaction models.TransactionJoinGold
		if err := tr.gormDb.ScanRows(rows, &transaction.Transaction); err != nil {
			return transactions, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}