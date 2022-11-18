package repositories

import (
	"time"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/infrastructure/database"
)

func (gr *goldRepository) NewGoldInventory(goldDetailID uint32, quantity int) error {
	insertGoldInventorySql := `INSERT INTO gold_inventories (gold_inventory_id, gold_detail_id, status, date_in, date_sold) VALUES ($1, $2, $3, $4, $5);`
	for i := 0; i < quantity; i++{
		id, err := database.GenerateUUID()
		if err != nil {
			return err
		}
		loc, _ := time.LoadLocation("Asia/Jakarta")
		if _, err := gr.db.Exec(gr.ctx, insertGoldInventorySql, id, goldDetailID, "safe", time.Now().In(loc), time.Time{}); err != nil {
			return err
		}
	}
	return nil
}

func (gr *goldRepository) UpdateGoldInventoryStatus(goldInventoryID uint32, status string) error {
	updateGoldInventoryStatus := `UPDATE gold_inventories SET status = $1 WHERE gold_inventory_id = $2`
	s := "safe"
	if status == "safe" {
		s = "front"
	}
	_, err := gr.db.Exec(gr.ctx, updateGoldInventoryStatus, s, goldInventoryID)
	return err
}