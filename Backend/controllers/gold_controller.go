package controllers

import (
	"net/http"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
	"github.com/gin-gonic/gin"
)

type goldController struct{
	goldUseCase domains.GoldUseCase
}

func NewGoldController(gu domains.GoldUseCase) *goldController{
	return &goldController{
		goldUseCase: gu,
	}
}

func (gc *goldController) NewGold(c *gin.Context){
	newGold := new(models.GoldDetail)
	if err := c.BindJSON(newGold); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := gc.goldUseCase.NewGold(newGold); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}