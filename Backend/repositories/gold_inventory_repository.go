package repositories

import (
	"context"
	"time"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/infrastructure/database"
)

func (gr *goldRepository) NewGoldInventory(goldDetailID uint32) error{
	id, err := database.GenerateUUID()
	if err != nil {
		return err
	}
	insertGoldInventorySql := `INSERT INTO gold_inventories (id, gold_detail_id, status, date_in) VALUES ($1, $2, $3, $4);`
	loc, _ := time.LoadLocation("Asia/Jakarta")
	ctx := context.Background()
	if _, err := gr.db.Exec(ctx, insertGoldInventorySql, id, goldDetailID, "s", time.Now().In(loc)); err != nil {
		return err
	}
	return nil
}