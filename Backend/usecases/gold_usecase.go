package usecases

import (
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
)

type goldUseCase struct{
	goldRepo domains.GoldRepository
}

func NewGoldUseCase(gr domains.GoldRepository) domains.GoldUseCase{
	return &goldUseCase{gr}
}

func (gr *goldUseCase) NewGold(goldDetail *models.GoldDetail) error{
	id, err := gr.goldRepo.NewGoldDetail(goldDetail)
	if err != nil {
		return err
	}
	err = gr.goldRepo.NewGoldInventory(id)
	if err != nil {
		return err
	}
	return nil
}