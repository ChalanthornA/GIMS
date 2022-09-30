package users

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type RegisterAdminBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Secret   string `json:"secret"`
}

func (h handler) RegisterAdmin(c *fiber.Ctx) error {
	realSecret := "BuBuBaBa" //เดี๋ยวยัดใส่ env อีกที
	newAdmin := new(RegisterAdminBody)
	if err := c.BodyParser(newAdmin); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Body")
	}
	if newAdmin.Secret != realSecret {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid secret")
	}

	if err := newAdmin.InsertAdmin(h.DB); err != nil{
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
	})
}

func (rab RegisterAdminBody) InsertAdmin(db *pgxpool.Pool) error{
	password := []byte(rab.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	ctx := context.Background()
	insertAdminSql := `INSERT INTO users (username, password, role) VALUES ($1, $2, 'admin')`
	_, err = db.Exec(ctx, insertAdminSql, rab.Username, string(hashedPassword))
	if err != nil {
		return err
	}
	return nil
}
