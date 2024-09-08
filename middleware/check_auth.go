package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func CheckAuth(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	fmt.Println(token)

	return c.Next()
}
