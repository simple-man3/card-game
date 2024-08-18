package controller

import (
	"card-game/models"
	"card-game/responses"
	"card-game/services"
	"card-game/validator"
	"github.com/gofiber/fiber/v2"
)

type Day int

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return responses.BodyParseErrToResponse(c)
	}

	if errs := validator.Validator.Struct(user); errs != nil {
		return responses.ValidationErrToResponse(errs, c)
	}

	if err := services.CreateUser(&user); err != nil {
		return responses.ServiceErrorToResponse(err, c)
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
