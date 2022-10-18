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

func (gu *goldUseCase) NewGold(goldDetail *models.GoldDetail) error{
	if err := gu.goldRepo.CheckGoldDetail(goldDetail); err != nil{
		return err
	}
	id, err := gu.goldRepo.NewGoldDetail(goldDetail)
	if err != nil {
		return err
	}
	err = gu.goldRepo.NewGoldInventory(id)
	if err != nil {
		return err
	}
	return nil
}

func (gu *goldUseCase) FindGoldDetailByCode(code string) ([]models.GoldDetail, error){
	details, err := gu.goldRepo.QueryGoldDetailByCode(code)
	return details, err
}