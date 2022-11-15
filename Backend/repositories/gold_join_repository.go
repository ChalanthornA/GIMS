package repositories

import "github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"

func (gr *goldRepository) QueryAllGoldDetailJoinInventory() ([]models.GoldDetailJoinInventory, error) {
	var golds []models.GoldDetailJoinInventory
	queryAllGoldDetailSql := `SELECT * FROM gold_details;`
	queryGoldInventoryByGoldDetailIDSql := `SELECT * FROM gold_inventories WHERE gold_detail_id = $1;`
	rows, err := gr.db.Query(gr.ctx, queryAllGoldDetailSql)
	if err != nil {
		return golds, err
	}
	for rows.Next() {
		var gold models.GoldDetailJoinInventory
		if err = rows.Scan(&gold.GoldDetailID, &gold.Code, &gold.Type, &gold.Detail, &gold.Weight, &gold.GoldPercent, &gold.GoldSmithFee, &gold.Picture, &gold.OtherDetail); err != nil {
			return golds, err
		}
		//query gold_inventory for append to array
		var inventories []models.GoldInventory
		rowsGoldInventories, err2 := gr.db.Query(gr.ctx, queryGoldInventoryByGoldDetailIDSql, gold.GoldDetailID)
		if err2 != nil {
			return golds, err2
		}
		for rowsGoldInventories.Next() {
			var inventory models.GoldInventory
			if err = rowsGoldInventories.Scan(&inventory.GoldInventoryID, &inventory.GoldDetailID, &inventory.Status, &inventory.DateIn); err != nil {
				return golds, err
			}
			inventories = append(inventories, inventory)
		}
		gold.Inventories = inventories
		golds = append(golds, gold)
	}
	return golds, nil
}