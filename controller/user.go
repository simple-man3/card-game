package controller

import (
	"card-game/models"
	"card-game/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	result := services.ExistUser(models.User{
		Name: "text",
	})
	fmt.Println(result)

	return c.SendString("create user")
}
