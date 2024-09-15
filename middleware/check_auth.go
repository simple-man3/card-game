package middleware

import (
	"card-game/responses"
	"card-game/services"
	"card-game/session"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func CheckAuth(c *fiber.Ctx) error {
	sess, _ := session.Store.Get(c)

	token := c.Get("Authorization")
	sessToken := sess.Get("Authorization")

	if token == "" && sessToken == nil {
		return responses.ForbiddenResponse(c)
	}

	authService := services.NewAuthService()
	if token != "" {
		authService.VerifyToken(token)
	}

	fmt.Println(token)
	fmt.Println(token)

	return c.Next()
}
