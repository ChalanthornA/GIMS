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
	createTableUser(dbpool, ctx)
	return dbpool
}

func createTableUser(dbpool *pgxpool.Pool, ctx context.Context){
	createTableUserSql := `CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, username VARCHAR(50), password VARCHAR(100), role VARCHAR(50));`
	_, err := dbpool.Exec(ctx, createTableUserSql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create SENSORS table: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully created table users")
}
