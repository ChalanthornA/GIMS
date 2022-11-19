package usecases

import (
	"strconv"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
)

type goldUseCase struct {
	goldRepo domains.GoldRepository
}

func NewGoldUseCase(gr domains.GoldRepository) domains.GoldUseCase {
	return &goldUseCase{gr}
}

func (gu *goldUseCase) ConvertIDStringToUint32(id string) (uint32, error) {
	goldDetailID, err := strconv.ParseUint(id, 10, 32)
	return uint32(goldDetailID), err
}

func (gu *goldUseCase) NewGold(goldDetail *models.GoldDetail) error {
	if err := gu.goldRepo.CheckGoldDetail(goldDetail); err != nil {
		return err
	}
	id, err := gu.goldRepo.NewGoldDetail(goldDetail)
	if err != nil {
		return err
	}
	newGoldInventory := &models.InputNewGoldInventory{GoldDetailID: id, Quantity: 1}
	err = gu.goldRepo.NewGoldInventory(newGoldInventory)
	if err != nil {
		return err
	}
	return nil
}

func (gu *goldUseCase) AddGold(newGoldInventory *models.InputNewGoldInventory) error {
	err := gu.goldRepo.NewGoldInventory(newGoldInventory)
	return err
}

func (gu *goldUseCase) GetAllGoldDetail() ([]models.GoldDetail, error) {
	res, err := gu.goldRepo.QueryAllGoldDetail()
	return res, err
}

func (gu *goldUseCase) FindGoldDetailByCode(code string) ([]models.GoldDetail, error) {
	details, err := gu.goldRepo.QueryGoldDetailByCode(code)
	return details, err
}

func (gu *goldUseCase) FindGoldDetailByDetail(g *models.GoldDetail) ([]models.GoldDetail, error) {
	details, err := gu.goldRepo.QueryGoldDetailByDetail(g)
	return details, err
}

func (gu *goldUseCase) EditGoldDetail(goldDetail *models.GoldDetail) error{
	return gu.goldRepo.UpdateGoldDetail(goldDetail)
}

func (gu *goldUseCase) GetAllGoldDetailJoinInventory() ([]models.GoldDetailJoinInventory, error) {
	return gu.goldRepo.QueryAllGoldDetailJoinInventory()
}

func (gu *goldUseCase) SetStatusGoldDetailToDelete(goldDetailID uint32) error {
	return gu.goldRepo.SetStatusGoldDetail(goldDetailID, "delete")
}

func (gu *goldUseCase) SetStatusGoldDetailToNormal(goldDetailID uint32) error {
	return gu.goldRepo.SetStatusGoldDetail(goldDetailID, "normal")
}