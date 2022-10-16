package domains

import "github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"

type GoldUseCase interface{
	NewGold(goldDetail *models.GoldDetail) error
}

type GoldRepository interface{
	NewGoldDetail(g *models.GoldDetail) (uint32, error)
	NewGoldInventory(goldDetailID uint32) error
}