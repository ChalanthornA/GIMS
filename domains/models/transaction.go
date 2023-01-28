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
	Username        string    `json:"username"`   //username
	BuyPrice        float64   `json:"buy_price"`  //use when TransactionType is change
	SellPrice       float64   `json:"sell_price"` //use when TransactionType is change
	Note            string    `json:"note"`
}

func (t *Transaction) SetTimeNow() {
	t.Date = time.Now()
}

type TransactionJoinGold struct {
	Transaction   Transaction   `json:"transaction"`
	GoldDetail    GoldDetail    `json:"gold_detail"`
	GoldInventory GoldInventory `json:"gold_inventory"`
}

type Report struct {
	Transactions []TransactionJoinGold `json:"transactions"`
	TotalPrice float64 `json:"total_price"`
	IncomePrice float64 `json:"income_price"`
	OutcomePrice float64 `json:"outcome_price"`
}