package models

import "time"

type GoldInventory struct {
	GoldInventoryID uint32    `json:"gold_inventory_id"`
	GoldDetailID    uint32    `json:"gold_detail_id"`
	Status          string    `json:"status"` //safe or front or sold
	DateIn          time.Time `json:"date_in"`
	DateSold        time.Time `json:"date_sold"`
	Note            string    `json:"note"`
}
