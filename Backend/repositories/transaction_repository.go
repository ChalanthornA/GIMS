package repositories

import (
	"context"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/infrastructure/database"
	"github.com/jackc/pgx/v4/pgxpool"
)

type transactionRepository struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func NewTransactionRepository(db *pgxpool.Pool) domains.TransactionRepository {
	return &transactionRepository{db: db, ctx: context.Background()}
}

func (tr *transactionRepository) InsertNewTransaction(transaction *models.Transaction) error{
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