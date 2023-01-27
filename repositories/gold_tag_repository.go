package repositories

import "github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"

func (gr *goldRepository) QueryGoldByTagSerialNumber(serialNumber uint32) *models.GoldInventory {
	var goldInventory models.GoldInventory
	queryGoldByTagSerialNumberSql := `SELECT * FROM gold_inventories WHERE tag_serial_number = ? AND is_sold = ?;`
	gr.gormDb.Raw(queryGoldByTagSerialNumberSql, serialNumber, 0).Scan(&goldInventory)
	return &goldInventory
}

func (gr *goldRepository) SetTagSerialNumberGoldInventory(id, serialNumber uint32) error {
	updateTagSerialNumberGoldInventorySql := `UPDATE gold_inventories SET tag_serial_number = $1 WHERE gold_inventory_id = $2 and is_sold = $3;`
	_, err := gr.db.Exec(gr.ctx, updateTagSerialNumberGoldInventorySql, serialNumber, id, 0)
	return err
}