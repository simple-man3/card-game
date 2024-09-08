package middleware

import (
	"card-game/session"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func CheckAuth(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	session, _ := session.Session.Get(c)

	token = session.Get("Dsada")

	fmt.Println(token)
	fmt.Println(token)

	return c.Next()
}
