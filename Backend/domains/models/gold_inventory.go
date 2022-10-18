package models

import "time"

type GoldInventory struct {
	GoldInventoryID uint32    `json:"gold_inventory_id"`
	GoldDetailID    uint32    `json:"gold_detail_id"`
	Status          string    `json:"status"`
	DateIn          time.Time `json:"date_in"`
}
