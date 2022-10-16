package models

type GoldDetail struct {
	ID           uint32  `json:"id"`
	Code         string  `json:"code"`
	Type         string  `json:"type"`
	Detail       string  `json:"detail"`
	Weight       float64 `json:"weight"` //น่าจะหน่วยเป็นกรัมไปก่อน
	GoldPercent  float64 `json:"gold_percent"`
	GoldSmithFee float64 `json:"gold_smith_fee"`
	Picture      string  `json:"picture"`
	OtherDetail  string  `json:"other_detail"`
}
