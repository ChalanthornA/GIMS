package models

type TransactionJoinGold struct {
	Transaction   Transaction   `json:"transaction"`
	GoldDetail    GoldDetail    `json:"gold_detail"`
	GoldInventory GoldInventory `json:"gold_inventory"`
}

type Report struct {
	Transactions       []TransactionJoinGold `json:"transactions"`
	TotalPrice         float64               `json:"total_price"`
	IncomePrice        float64               `json:"income_price"`
	OutcomePrice       float64               `json:"outcome_price"`
	BuyPrice           float64               `json:"buy_price"`  //sum of all transaction type buy
	SellPrice          float64               `json:"sell_price"` //sum of all transaction type sell
	ChangeIncomePrice  float64               `json:"change_income_price"`
	ChangeOutcomePrice float64               `json:"change_outcome_price"`
	TotalChangePrice   float64               `json:"total_change_price"`
}