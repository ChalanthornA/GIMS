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

func (gu *goldUseCase) NewGold(g *models.InputNewGoldDetail) error {
	newGoldDetail := &models.GoldDetail{Code: g.Code, Type: g.Type, Detail: g.Detail, Weight: g.Weight, GoldPercent: g.GoldPercent, GoldSmithFee: g.GoldSmithFee, Picture: g.Picture, Status: g.Status}
	if err := gu.goldRepo.CheckGoldDetail(newGoldDetail); err != nil {
		return err
	}
	id, err := gu.goldRepo.NewGoldDetail(newGoldDetail)
	if err != nil {
		return err
	}
	newGoldInventory := &models.InputNewGoldInventory{GoldDetailID: id, Note: g.Note, Quantity: g.Quantity}
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

func (gu *goldUseCase) EditGoldDetail(goldDetail *models.GoldDetail) error {
	return gu.goldRepo.UpdateGoldDetail(goldDetail)
}

func (gu *goldUseCase) GetAllGoldDetailJoinInventory() ([]models.GoldDetailJoinInventory, error) {
	return gu.goldRepo.QueryAllGoldDetailJoinInventory()
}

func (gu *goldUseCase) SetStatusGoldDetailToDelete(goldDetailID []uint32) error {
	for _, id := range goldDetailID {
		if err := gu.goldRepo.SetStatusGoldDetail(id, "delete"); err != nil {
			return err
		}
	}
	return nil
}

func (gu *goldUseCase) SetStatusGoldDetailToNormal(goldDetailID uint32) error {
	return gu.goldRepo.SetStatusGoldDetail(goldDetailID, "normal")
}

func (gu *goldUseCase) SetStatusGoldInventory(goldInventoryIDs []uint32, status string) error {
	for _, id := range goldInventoryIDs{
		if err := gu.goldRepo.UpdateGoldInventoryStatus(id, status); err != nil {
			return err
		}
	}
	return nil
}

func (gu *goldUseCase) GetGoldDetailJoinInventoryByDetail(g *models.GoldDetail) ([]models.GoldDetailJoinInventory, error) {
	return gu.goldRepo.QueryGoldDetailJoinInventoryByDetail(g)
}

func (gu *goldUseCase) GetGoldDetailByGoldDetailID(goldDetailID uint32) (models.GoldDetail, error) {
	return gu.goldRepo.QueryGoldDetailByGoldDetailID(goldDetailID)
}

func (gu *goldUseCase) DeleteGoldInventoryByIDArray(ids []uint32) error {
	for _, id := range ids {
		if err := gu.goldRepo.DeleteGoldInventoryByID(id); err != nil {
			return err
		}
	}
	return nil
}

func (gu *goldUseCase) SetTagSerialNumberGoldInventory(input *models.InputSetTagSerialNumber) error {
	return gu.goldRepo.SetTagSerialNumberGoldInventory(input.GoldInventoryID, input.TagSerialNumber)
}