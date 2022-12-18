package repositories

import (
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
)

func (gr *goldRepository) QueryAllGoldDetailJoinInventory() ([]models.GoldDetailJoinInventory, error) {
	var golds []models.GoldDetailJoinInventory
	queryAllGoldDetailSql := `SELECT * FROM gold_details WHERE status = $1;`
	queryGoldInventoryByGoldDetailIDSql := `SELECT * FROM gold_inventories WHERE gold_detail_id = $1 AND is_sold = $2;`
	rows, err := gr.db.Query(gr.ctx, queryAllGoldDetailSql, "normal")
	if err != nil {
		return golds, err
	}
	for rows.Next() {
		var gold models.GoldDetailJoinInventory
		if err = rows.Scan(&gold.GoldDetailID, &gold.Code, &gold.Type, &gold.Detail, &gold.Weight, &gold.GoldPercent, &gold.GoldSmithFee, &gold.Picture, &gold.Status); err != nil {
			return golds, err
		}
		var inventories []models.GoldInventory
		rowsGoldInventories, err2 := gr.db.Query(gr.ctx, queryGoldInventoryByGoldDetailIDSql, gold.GoldDetailID, 0)
		if err2 != nil {
			return golds, err2
		}
		for rowsGoldInventories.Next() {
			var inventory models.GoldInventory
			if err = rowsGoldInventories.Scan(&inventory.GoldInventoryID, &inventory.GoldDetailID, &inventory.Status, &inventory.DateIn, &inventory.DateSold, &inventory.Note, &inventory.IsSold); err != nil {
				return golds, err
			}
			inventories = append(inventories, inventory)
		}
		gold.Inventories = inventories
		golds = append(golds, gold)
	}
	return golds, nil
}