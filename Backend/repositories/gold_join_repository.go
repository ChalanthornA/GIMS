package repositories

import (
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
)

func (gr *goldRepository) QueryAllGoldDetailJoinInventory() ([]models.GoldDetailJoinInventory, error) {
	var golds []models.GoldDetailJoinInventory
	queryAllGoldDetailSql := `SELECT * FROM gold_details WHERE status = $1;`
	rows, err := gr.db.Query(gr.ctx, queryAllGoldDetailSql, "normal")
	if err != nil {
		return golds, err
	}
	for rows.Next() {
		var gold models.GoldDetailJoinInventory
		if err := rows.Scan(&gold.GoldDetailID, &gold.Code, &gold.Type, &gold.Detail, &gold.Weight, &gold.GoldPercent, &gold.GoldSmithFee, &gold.Picture, &gold.Status); err != nil {
			return golds, err
		}
		inventories, err := gr.QueryAllGoldInventoryByGoldDetailID(gold.GoldDetailID)
		if err != nil {
			return golds, err
		}
		gold.Inventories = inventories
		golds = append(golds, gold)
	}
	return golds, nil
}

func (gr *goldRepository) QueryGoldDetailJoinInventoryByDetail(g *models.GoldDetail) ([]models.GoldDetailJoinInventory, error) {
	var golds []models.GoldDetailJoinInventory
	queryAllGoldDetailSql := `SELECT * FROM gold_details WHERE status = $1;`
	rows, err := gr.db.Query(gr.ctx, queryAllGoldDetailSql, "normal")
	if err != nil {
		return golds, err
	}
	for rows.Next() {
		var gold models.GoldDetailJoinInventory
		if err := rows.Scan(&gold.GoldDetailID, &gold.Code, &gold.Type, &gold.Detail, &gold.Weight, &gold.GoldPercent, &gold.GoldSmithFee, &gold.Picture, &gold.Status); err != nil {
			return golds, err
		}
		if !checkGoldDetailJoinInventory(g, &gold) {
			continue
		}
		inventories, err := gr.QueryAllGoldInventoryByGoldDetailID(gold.GoldDetailID)
		if err != nil {
			return golds, err
		}
		gold.Inventories = inventories
		golds = append(golds, gold)
	}
	return golds, nil
}

func checkGoldDetailJoinInventory(g *models.GoldDetail, goldDetail *models.GoldDetailJoinInventory) bool {
	if g.Code != "" {
		if g.Code != goldDetail.Code {
			return false
		}
	}
	if g.Type != "" {
		if g.Type != goldDetail.Type {
			return false
		}
	}
	if g.Weight != 0 {
		if g.Weight != goldDetail.Weight {
			return false
		}
	}
	if g.GoldPercent != 0 {
		if g.GoldPercent != goldDetail.GoldPercent {
			return false
		}
	}
	if g.GoldSmithFee != 0 {
		if g.GoldSmithFee != goldDetail.GoldSmithFee {
			return false
		}
	}
	if g.Detail != "" {
		if g.Detail != goldDetail.Detail {
			return false
		}
	}
	return true
}