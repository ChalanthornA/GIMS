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
	if transaction.TransactionType != "buy" || transaction.GoldInventoryID != 0 || transaction.FromNote != "" || transaction.ToNote != "" {
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
	if transaction.TransactionType != "sell" || transaction.FromNote != "" || transaction.ToNote != "" {
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