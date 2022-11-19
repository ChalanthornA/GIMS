package domains

import "github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"

type GoldUseCase interface{
	NewGold(goldDetail *models.GoldDetail) error
	ConvertIDStringToUint32(id string) (uint32, error)
	AddGold(newGoldInventory *models.InputNewGoldInventory) error
	GetAllGoldDetail() ([]models.GoldDetail, error)
	FindGoldDetailByCode(code string) ([]models.GoldDetail, error)
	FindGoldDetailByDetail(g *models.GoldDetail) ([]models.GoldDetail, error)
	EditGoldDetail(goldDetail *models.GoldDetail) error
	GetAllGoldDetailJoinInventory() ([]models.GoldDetailJoinInventory, error)
	SetStatusGoldDetailToDelete(goldDetailID uint32) error
	SetStatusGoldDetailToNormal(goldDetailID uint32) error
}

type GoldRepository interface{
	NewGoldDetail(g *models.GoldDetail) (uint32, error)
	NewGoldInventory(newGoldInventory *models.InputNewGoldInventory) error
	QueryAllGoldDetail() ([]models.GoldDetail, error)
	QueryGoldDetailByCode(code string) ([]models.GoldDetail, error)
	CheckGoldDetail(g *models.GoldDetail) error
	QueryGoldDetailByDetail(g *models.GoldDetail) ([]models.GoldDetail, error)
	UpdateGoldDetail(goldDetail *models.GoldDetail) error
	QueryAllGoldDetailJoinInventory() ([]models.GoldDetailJoinInventory, error)
	SetStatusGoldDetail(goldDetailID uint32, setStatus string) error
}