package models

import "time"

type Transaction struct {
	TransactionID   uint32    `json:"transaction_id"`
	TransactionType string    `json:"transaction_type"` //buy sell change
	Date            time.Time `json:"date"`
	GoldPrice       string    `json:"gold_price"` //format is 28xxx - 29xxx <- just example
	Weight          float64   `json:"weight"`
	Price           float64   `json:"price"`
	GoldDetailID    uint32    `json:"gold_detail_id"`
	GoldInventoryID uint32    `json:"gold_inventory_id"`
	Username        string    `json:"username"`  //username
	FromNote        string    `json:"from_note"` //use when TransactionType is change
	ToNote          string    `json:"to_note"`   //use when TransactionType is change
	Note            string    `json:"note"`
}

func (t *Transaction) SetTimeNow() {
	t.Date = time.Now()
}
