package controllers

import (
	"net/http"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
	"github.com/gin-gonic/gin"
)

type goldController struct {
	goldUseCase  domains.GoldUseCase
}

func NewGoldController(gu domains.GoldUseCase) *goldController {
	return &goldController{
		goldUseCase:  gu,
	}
}

func (gc *goldController) NewGold(c *gin.Context) {
	newGold := new(models.InputNewGoldDetail)
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

func (gc *goldController) AddGold(c *gin.Context) {
	newGoldInventory := new(models.InputNewGoldInventory)
	if err := c.BindJSON(newGoldInventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := gc.goldUseCase.AddGold(newGoldInventory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (gc *goldController) GetAllGoldDetail(c *gin.Context) {
	goldDetails, err := gc.goldUseCase.GetAllGoldDetail()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": goldDetails,
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

func (gc *goldController) EditGoldDetail(c *gin.Context) {
	goldDetail := new(models.GoldDetail)
	if err := c.BindJSON(goldDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := gc.goldUseCase.EditGoldDetail(goldDetail); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (gc *goldController) GetAllGoldDetailJoinInventory(c *gin.Context) {
	res, err := gc.goldUseCase.GetAllGoldDetailJoinInventory()
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

func (gc *goldController) SetStatusGoldDetailToDelete(c *gin.Context) {
	idString := c.Param("id")
	idUint32, err := gc.goldUseCase.ConvertIDStringToUint32(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := gc.goldUseCase.SetStatusGoldDetailToDelete(idUint32); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (gc *goldController) SetStatusGoldDetailToNormal(c *gin.Context) {
	idString := c.Param("id")
	idUint32, err := gc.goldUseCase.ConvertIDStringToUint32(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := gc.goldUseCase.SetStatusGoldDetailToNormal(idUint32); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (gc *goldController) SetStatusGoldInventory(c *gin.Context) {
	input := new(models.InputSetStatusGoldInventory)
	if err := c.BindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := gc.goldUseCase.SetStatusGoldInventory(input.GoldInventoryID, input.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (gc *goldController) GetGoldDetailJoinInventoryByDetail(c *gin.Context) {
	goldDetail := new(models.GoldDetail)
	if err := c.BindJSON(goldDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	res, err := gc.goldUseCase.GetGoldDetailJoinInventoryByDetail(goldDetail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data": res,
	})
}