package repositories

import (
	"context"
	"fmt"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/infrastructure/database"
	"github.com/jackc/pgx/v4/pgxpool"
)

type goldRepository struct {
	ctx context.Context
	db  *pgxpool.Pool
}

func NewGoldRepository(db *pgxpool.Pool) domains.GoldRepository {
	return &goldRepository{ctx: context.Background(), db: db}
}

func (gr *goldRepository) NewGoldDetail(g *models.GoldDetail) (uint32, error) {
	id, err := database.GenerateUUID()
	if err != nil {
		return 0, err
	}
	g.GoldDetailID = id
	insertGoldDetailSql := `INSERT INTO gold_details (gold_detail_id, code, type, detail, weight, gold_percent, gold_smith_fee, picture, other_detail) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`
	if _, err := gr.db.Exec(gr.ctx, insertGoldDetailSql, g.GoldDetailID, g.Code, g.Type, g.Detail, g.Weight, g.GoldPercent, g.GoldSmithFee, g.Picture, g.OtherDetail); err != nil {
		return 0, err
	}
	return id, nil
}

func (gr *goldRepository) CheckGoldDetail(g *models.GoldDetail) error {
	queryGoldDetailByDetail := `
		SELECT * 
		FROM gold_details
		WHERE code = $1 AND type = $2 AND detail = $3 AND weight = $4 AND gold_percent = $5 AND gold_smith_fee = $6 AND other_detail = $7;
	`
	rows, err := gr.db.Query(gr.ctx, queryGoldDetailByDetail, g.Code, g.Type, g.Detail, g.Weight, g.GoldPercent, g.GoldSmithFee, g.OtherDetail)
	if err != nil {
		return err
	}
	detail := new(models.GoldDetail)
	for rows.Next() {
		if err = rows.Scan(&detail.GoldDetailID, &detail.Code, &detail.Type, &detail.Detail, &detail.Weight, &detail.GoldPercent, &detail.GoldSmithFee, &detail.Picture, &detail.OtherDetail); err != nil {
			return err
		}
	}
	if detail.Type != "" {
		return fmt.Errorf("this detail is already exist see code %s", detail.Code)
	}
	return nil
}

func (gr *goldRepository) QueryAllGoldDetail() ([]models.GoldDetail, error) {
	var res []models.GoldDetail
	queryAllGoldDetail := `SELECT * FROM gold_details;`
	rows, err := gr.db.Query(gr.ctx, queryAllGoldDetail)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		var detail models.GoldDetail
		if err = rows.Scan(&detail.GoldDetailID, &detail.Code, &detail.Type, &detail.Detail, &detail.Weight, &detail.GoldPercent, &detail.GoldSmithFee, &detail.Picture, &detail.OtherDetail); err != nil {
			return res, err
		}
		res = append(res, detail)
	}
	return res, nil
}

func (gr *goldRepository) QueryGoldDetailByCode(code string) ([]models.GoldDetail, error) {
	var res []models.GoldDetail
	queryGoldDetailByCodeSql := `
		SELECT *
		FROM gold_details
		WHERE code = $1;
	`
	rows, err := gr.db.Query(gr.ctx, queryGoldDetailByCodeSql, code)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		var detail models.GoldDetail
		if err = rows.Scan(&detail.GoldDetailID, &detail.Code, &detail.Type, &detail.Detail, &detail.Weight, &detail.GoldPercent, &detail.GoldSmithFee, &detail.Picture, &detail.OtherDetail); err != nil {
			return res, err
		}
		res = append(res, detail)
	}
	return res, nil
}

func (gr *goldRepository) QueryGoldDetailByDetail(g *models.GoldDetail) ([]models.GoldDetail, error) {
	var res []models.GoldDetail
	queryAllGoldDetailSql := `SELECT * FROM gold_details;`
	rows, err := gr.db.Query(gr.ctx, queryAllGoldDetailSql)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		var goldDetail models.GoldDetail
		if err = rows.Scan(&goldDetail.GoldDetailID, &goldDetail.Code, &goldDetail.Type, &goldDetail.Detail, &goldDetail.Weight, &goldDetail.GoldPercent, &goldDetail.GoldSmithFee, &goldDetail.Picture, &goldDetail.OtherDetail); err != nil {
			return res, err
		}
		if g.Type != "" {
			if g.Type != goldDetail.Type {
				continue
			}
		}
		if g.Weight != 0 {
			if g.Weight != goldDetail.Weight {
				continue
			}
		}
		if g.GoldPercent != 0 {
			if g.GoldPercent != goldDetail.GoldPercent {
				continue
			}
		}
		if g.GoldSmithFee != 0 {
			if g.GoldSmithFee != goldDetail.GoldSmithFee {
				continue
			}
		}
		if g.OtherDetail != "" {
			if g.OtherDetail != goldDetail.OtherDetail {
				continue
			}
		}
		if g.Detail != "" {
			if g.Detail != goldDetail.Detail {
				continue
			}
		}
		res = append(res, goldDetail)
	}
	return res, nil
}
