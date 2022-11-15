package repositories

import (
	"time"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/infrastructure/database"
)

func (gr *goldRepository) NewGoldInventory(goldDetailID uint32) error {
	id, err := database.GenerateUUID()
	if err != nil {
		return err
	}
	insertGoldInventorySql := `INSERT INTO gold_inventories (gold_inventory_id, gold_detail_id, status, date_in) VALUES ($1, $2, $3, $4);`
	loc, _ := time.LoadLocation("Asia/Jakarta")
	if _, err := gr.db.Exec(gr.ctx, insertGoldInventorySql, id, goldDetailID, "s", time.Now().In(loc)); err != nil {
		return err
	}
	return nil
}

func (gr *goldRepository) UpdateGoldInventoryStatus(goldInventoryID uint32, status string) error {
	updateGoldInventoryStatus := `UPDATE gold_inventories SET status = $1 WHERE gold_inventory_id = $2`
	s := "s"
	if status == "s" {
		s = "f"
	}
	_, err := gr.db.Exec(gr.ctx, updateGoldInventoryStatus, s, goldInventoryID)
	return err
}