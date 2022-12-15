package repositories

import (
	"fmt"
	"time"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/infrastructure/database"
)

func (gr *goldRepository) NewGoldInventory(newGoldInventory *models.InputNewGoldInventory) error {
	insertGoldInventorySql := `INSERT INTO gold_inventories (gold_inventory_id, gold_detail_id, status, date_in, date_sold, note) VALUES ($1, $2, $3, $4, $5, $6);`
	for i := 0; i < newGoldInventory.Quantity; i++ {
		loc, _ := time.LoadLocation("Asia/Jakarta")
		if _, err := gr.db.Exec(gr.ctx, insertGoldInventorySql, database.GenerateUUID(), newGoldInventory.GoldDetailID, "safe", time.Now().In(loc), time.Time{}, newGoldInventory.Note); err != nil {
			return err
		}
	}
	return nil
}

func (gr *goldRepository) UpdateGoldInventoryStatus(goldInventoryID uint32, status string) error {
	updateGoldInventoryStatus := `UPDATE gold_inventories SET status = $1 WHERE gold_inventory_id = $2`
	_, err := gr.db.Exec(gr.ctx, updateGoldInventoryStatus, status, goldInventoryID)
	return err
}

func (gr *goldRepository) CheckGoldInventoryByGoldInventoryID(id uint32) (models.GoldInventory, error) {
	var goldInventory models.GoldInventory
	queryGoldInventoryByGoldInventoryIDSql := `SELECT * FROM gold_inventories WHERE gold_inventory_id = $1 AND status != $2;`
	rows, err := gr.db.Query(gr.ctx, queryGoldInventoryByGoldInventoryIDSql, id, "sold")
	if err != nil {
		return goldInventory, err
	}
	for rows.Next() {
		if err = rows.Scan(&goldInventory.GoldInventoryID, &goldInventory.GoldDetailID, &goldInventory.Status, &goldInventory.DateIn, &goldInventory.DateSold, &goldInventory.Note); err != nil {
			return goldInventory, err
		}
	}
	if goldInventory.GoldInventoryID == 0 {
		return goldInventory, fmt.Errorf("product not found")
	}
	return goldInventory, err
}