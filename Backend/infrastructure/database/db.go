package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func NewDb() *pgxpool.Pool{
	ctx := context.Background()
	connStr := "postgres://postgres:ppaallmm@localhost:5432/gims"
	dbpool, err := pgxpool.Connect(ctx, connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	inItTable(dbpool, ctx)
	return dbpool
}

func inItTable(dbpool *pgxpool.Pool, ctx context.Context){
	createUserTableSql := `CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, username VARCHAR(50) UNIQUE, password VARCHAR(100), role VARCHAR(50));`
	if _, err := dbpool.Exec(ctx, createUserTableSql); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create users table: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully created users table")
	createGoldDetailTableSql := `
		CREATE TABLE IF NOT EXISTS gold_details (
			gold_detail_id BIGINT NOT NULL, 
			code VARCHAR(50),
			type VARCHAR(50), 
			detail VARCHAR(100), 
			weight FLOAT,
			gold_percent FLOAT, 
			gold_smith_fee FLOAT, 
			picture VARCHAR(100), 
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
