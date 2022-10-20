package controllers

import (
	"net/http"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
	"github.com/gin-gonic/gin"
)

type goldController struct {
	goldUseCase domains.GoldUseCase
}

func NewGoldController(gu domains.GoldUseCase) *goldController {
	return &goldController{
		goldUseCase: gu,
	}
}

func (gc *goldController) NewGold(c *gin.Context) {
	role, _ := c.Get("role")
	if role != "admin" && role != "owner" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "only admin or owner can do this",
		})
		return
	}
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

func (gc *goldController) FindGoldDetailByCode(c *gin.Context) {
	code := c.Param("code")
	res, err := gc.goldUseCase.FindGoldDetailByCode(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (gc *goldController) FindGoldDetailByDetail(c *gin.Context) {
	goldDetail := new(models.GoldDetail)
	if err := c.BindJSON(goldDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	res, err := gc.goldUseCase.FindGoldDetailByDetail(goldDetail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
