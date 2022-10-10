package controllers

import (
	"net/http"

	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
	"github.com/gin-gonic/gin"
)

type registerAdminBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Secret   string `json:"secret"`
}

type UserController struct {
	userUseCase domains.UserUseCase
}

func NewUserController(uu domains.UserUseCase) *UserController {
	return &UserController{
		userUseCase: uu,
	}
}

func (uc *UserController) RegisterAdmin(c *gin.Context) {
	body := new(registerAdminBody)
	if err := c.BindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	newAdmin := &models.User{
		Username: body.Username,
		Password: body.Password,
	}
	if err := uc.userUseCase.RegisterAdmin(newAdmin, body.Secret); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (uc *UserController) SignIn(c *gin.Context){
	body := new(models.User)
	if err := c.BindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	token, err := uc.userUseCase.SignIn(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"accesstoken": token,
	})
}

func (uc *UserController) TestJWT(c *gin.Context){
	username, _ := c.Get("username")
	role, _ := c.Get("role")
	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"role": role,
	})
}