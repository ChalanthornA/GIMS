package repositories

import (
	"context"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/infrastructure/database"
	"github.com/jackc/pgx/v4/pgxpool"
)

type goldRepository struct{
	db *pgxpool.Pool
}

func NewGoldRepository(db *pgxpool.Pool) domains.GoldRepository{
	return &goldRepository{db}
}

func (gr *goldRepository) NewGoldDetail(g *models.GoldDetail) (uint32, error){
	id, err := database.GenerateUUID()
	if err != nil {
		return 0, err
	}
	g.ID = id
	insertGoldDetailSql := `INSERT INTO gold_details (id, code, type, detail, weight, gold_percent, gold_smith_fee, picture, other_detail) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`
	ctx := context.Background()
	if _, err := gr.db.Exec(ctx, insertGoldDetailSql, g.ID, g.Code, g.Type, g.Detail, g.Weight, g.GoldPercent, g.GoldSmithFee, g.Picture, g.OtherDetail); err != nil {
		return 0, err
	}
	return id, nil
}