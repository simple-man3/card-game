package controller

import (
	"card-game/requests"
	"card-game/responses"
	"card-game/services"
	"card-game/validator"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var request requests.LoginRequest

	if err := c.BodyParser(&request); err != nil {
		return responses.BodyParseErrToResponse()
	}

	if errs := validator.Validator.Struct(request); errs != nil {
		return responses.ValidationErrToResponse(errs, c)
	}

	token, err := services.Auth(c, request.Email, request.Password)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	return c.Status(fiber.StatusOK).JSON(responses.LoginResponse{Token: token})
}
