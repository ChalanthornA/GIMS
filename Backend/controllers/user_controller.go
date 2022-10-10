package controllers

import (
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains"
	"github.com/ChalanthornA/Gold-Inventory-Management-System/domains/models"
	"github.com/gofiber/fiber/v2"
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

func (uc *UserController) RegisterAdmin(c *fiber.Ctx) error {
	body := new(registerAdminBody)
	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Body")
	}
	newAdmin := &models.User{
		Username: body.Username,
		Password: body.Password,
	}
	if err := uc.userUseCase.RegisterAdmin(newAdmin, body.Secret); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
	})
}

func (uc *UserController) SignIn(c *fiber.Ctx) error{
	body := new(models.User)
	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Body")
	}
	token, err := uc.userUseCase.SignIn(body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"accesstoken": token,
	})
}

func (uc *UserController) TestJWT(c *fiber.Ctx) error{
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"username": c.Locals("username"),
		"role": c.Locals("role"),
	})
}