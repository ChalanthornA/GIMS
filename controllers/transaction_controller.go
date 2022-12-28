package controllers

import (
	"fmt"
	"net/http"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
	"github.com/gin-gonic/gin"
)

type transactionController struct {
	transactionUseCase domains.TransactionUseCase
}

func NewTransactionController(tu domains.TransactionUseCase) *transactionController{
	return &transactionController{transactionUseCase: tu}
}

func setUsername(c *gin.Context, t *models.Transaction) {
	username, _ := c.Get("username")
	t.Username = fmt.Sprint(username)
}

func (tc *transactionController) NewTransactionBuy(c *gin.Context) {
	transaction := new(models.Transaction)
	if err := c.BindJSON(transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if transaction.TransactionType != "buy" || transaction.GoldInventoryID != 0 || transaction.BuyPrice != 0 || transaction.SellPrice != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid transaction type buy body",
		})
		return
	}
	setUsername(c, transaction)
	if err := tc.transactionUseCase.NewTransactionTypeBuy(transaction); err != nil { 
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (tc *transactionController) NewTransactionSell(c *gin.Context) {
	transaction := new(models.Transaction)
	if err := c.BindJSON(transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if transaction.TransactionType != "sell" || transaction.BuyPrice != 0 || transaction.SellPrice != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid transaction type sell body",
		})
		return
	}
	setUsername(c, transaction)
	if err := tc.transactionUseCase.NewTransactionTypeSell(transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (tc *transactionController) NewTransactionChange(c *gin.Context) {
	transaction := new(models.Transaction)
	if err := c.BindJSON(transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if transaction.TransactionType != "change" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid transaction type change body",
		})
		return
	}
	setUsername(c, transaction)
	if err := tc.transactionUseCase.NewTransactionTypeChange(transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (tc *transactionController) RollBackTransaction(c *gin.Context) {
	transaction := new(models.Transaction)
	if err := c.BindJSON(transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := tc.transactionUseCase.RollBackTransaction(transaction.TransactionID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (tc *transactionController) GetAllTransactionJoinGold(c *gin.Context) {
	tjgs, err := tc.transactionUseCase.GetAllTransactionJoinGold()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data": tjgs,
	})
}

func (tc *transactionController) GetAllTransactionByTransactionType(c *gin.Context) {
	transactionType := c.Query("type")
	tjgs, err := tc.transactionUseCase.GetTransactionByTransactionType(transactionType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data": tjgs,
	})
}

func (tc *transactionController) GetTransactionByTimeInterval(c *gin.Context) {
	timeRange := c.Param("range")
	fmt.Println(timeRange)
	tjgs, err := tc.transactionUseCase.GetTransactionByTimeInterval(timeRange)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data": tjgs,
	})
}