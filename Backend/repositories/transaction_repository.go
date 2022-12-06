package repositories

import (
	"context"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains"
	"github.com/jackc/pgx/v4/pgxpool"
)

type transactionRepository struct {
	ctx context.Context
	db  *pgxpool.Pool
}

func NewTransactionRepository(db *pgxpool.Pool) domains.TransactionRepository {
	return &transactionRepository{ctx: context.Background(), db: db}
}