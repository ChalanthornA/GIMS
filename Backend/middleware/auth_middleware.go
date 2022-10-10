package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AuthorizationRequired() fiber.Handler {
    return jwtware.New(jwtware.Config{
        SuccessHandler: AuthSuccess,
        ErrorHandler:   AuthError,
        SigningKey:     []byte("secret"),
        SigningMethod: "HS256",
    })
}

func AuthError(c *fiber.Ctx, e error) error {
    c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
        "error": "Unauthorized",
        "msg":   e.Error(),
    })
    return nil
}

func AuthSuccess(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	c.Locals("username", claims["username"].(string))
	c.Locals("role", claims["role"].(string))
    c.Next()
    return nil
}