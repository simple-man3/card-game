package middleware

import (
	"card-game/responses"
	"card-game/services"
	"card-game/session"
	"github.com/gofiber/fiber/v2"
)

func CheckAuth(c *fiber.Ctx) error {
	sess, _ := session.Store.Get(c)

	headToken := c.Get("Authorization")
	sessToken := sess.Get("Authorization")

	if headToken == "" && sessToken == nil {
		return responses.ForbiddenResponse(c)
	}

	authService := services.NewAuthService()
	var token string

	if headToken != "" {
		token = headToken
	} else {
		token = sessToken.(string)
	}

	if err := authService.AuthFromToken(token); err != nil {
		return responses.ForbiddenResponse(c)
	}

	return c.Next()
}
