package controllers

import (
	"net/http"
	"strconv"

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

func (gc *goldController) AddGold(c *gin.Context) {
	id := c.Query("golddetailid")
	quantity := c.Query("quantity")
	goldDetailID, err := gc.goldUseCase.ConvertIDStringToUint32(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	intQuantity, err := strconv.Atoi(quantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := gc.goldUseCase.AddGold(goldDetailID, intQuantity); err != nil {
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