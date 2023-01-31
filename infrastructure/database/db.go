package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *pgxpool.Pool
var GormDB *gorm.DB

func NewDb() *pgxpool.Pool {
	ctx := context.Background()
	connStr := "postgres://postgres:ppaallmm@localhost:5432/gims"
	dbpool, err := pgxpool.Connect(ctx, connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to pgx database: %v\n", err)
		os.Exit(1)
	}
	createUserTableSql(dbpool, ctx)
	createGoldTable(dbpool, ctx)
	createTransactionTable(dbpool, ctx)
	return dbpool
}

func NewGormDb() *gorm.DB {
	dsn := "postgres://postgres:ppaallmm@localhost:5432/gims"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to gorm database: %v\n", err)
		os.Exit(1)
	}
	return db
}

func createUserTableSql(dbpool *pgxpool.Pool, ctx context.Context) {
	createUserTableSql := `CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, username VARCHAR(50) UNIQUE, password VARCHAR(100), role VARCHAR(50));`
	if _, err := dbpool.Exec(ctx, createUserTableSql); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create users table: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully created users table")
}

func createGoldTable(dbpool *pgxpool.Pool, ctx context.Context) {
	createGoldDetailTableSql := `
		CREATE TABLE IF NOT EXISTS gold_details (
			gold_detail_id BIGINT NOT NULL, 
			code VARCHAR(50),
			type VARCHAR(50), 
			detail VARCHAR(100), 
			weight FLOAT,
			gold_percent FLOAT, 
			gold_smith_fee FLOAT, 
			picture VARCHAR(500), 
			status VARCHAR(100),
			PRIMARY KEY (gold_detail_id)
		);
	`
	if _, err := dbpool.Exec(ctx, createGoldDetailTableSql); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create gold_details table: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully created gold_details table")
	createGoldInventoryTableSql := `
		CREATE TABLE IF NOT EXISTS gold_inventories (
			gold_inventory_id BIGINT NOT NULL,
			gold_detail_id BIGINT NOT NULL,
			status VARCHAR(50), 
			date_in TIMESTAMPTZ NOT NULL,
			date_sold TIMESTAMPTZ NOT NULL,
			note VARCHAR(300),
			is_sold INT,
			tag_serial_number BIGINT DEFAULT 0,
			FOREIGN KEY (gold_detail_id) REFERENCES gold_details(gold_detail_id)
		);
		SELECT create_hypertable('gold_inventories', 'date_sold', if_not_exists => TRUE);
	`
	if _, err := dbpool.Exec(ctx, createGoldInventoryTableSql); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create gold_details table: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully created gold_inventories table")
}

func createTransactionTable(dbpool *pgxpool.Pool, ctx context.Context) {
	createTransactionTableSql := `
		CREATE TABLE IF NOT EXISTS transactions (
			transaction_id BIGINT NOT NULL,
			transaction_type VARCHAR(50),
			date TIMESTAMPTZ NOT NULL,
			gold_price VARCHAR(100),
			weight FLOAT,
			price FLOAT,
			gold_detail_id BIGINT,
			gold_inventory_id BIGINT,
			username VARCHAR(100),
			buy_price FLOAT,
			sell_price FLOAT,
			note VARCHAR(300)
		);
		SELECT create_hypertable('transactions', 'date', if_not_exists => TRUE);
	`
	if _, err := dbpool.Exec(ctx, createTransactionTableSql); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create transactions table: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully created transactions table")
}