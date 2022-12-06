package models

import "time"

type Transaction struct {
	TransactionID   uint32    `json:"transaction_id"`
	TransactionType string    `json:"transaction_type"`
	Date            time.Time `json:"date"`
	GoldPrice       string    `json:"gold_price"` //format is 28xxx - 29xxx <- just example
	Weight          float64   `json:"weight"`
	Price           float64   `json:"price"`
	GoldDetailID    uint32    `json:"gold_detail_id"`
	User            string    `json:"user"` //username
	From            string    `json:"from"` //use when TransactionType is change
	To              string    `json:"to"`   //use when TransactionType is change
}
