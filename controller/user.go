package controller

import (
	"card-game/models"
	"card-game/services"
	"card-game/validator"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if errs := validator.Validator.Struct(user); errs != nil {
		return validator.GetValidationErrResponse(errs, c)
	}

	if err := services.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusBadRequest).JSON(user)
}
