package controller

import (
	"card-game/models"
	"card-game/requests"
	"card-game/responses"
	"card-game/services"
	"card-game/validator"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func CreateWallet(c *fiber.Ctx) error {
	var request requests.CreateWalletRequest

	if err := c.BodyParser(&request); err != nil {
		return responses.BodyParseErrToResponse()
	}

	if errs := validator.Validator.Struct(request); errs != nil {
		return responses.ValidationErrToResponse(errs, c)
	}

	bytes, err := json.Marshal(request)
	if err != nil {
		return responses.ServiceErrorToResponse(err)
	}

	var wallet *models.Wallet
	if err := json.Unmarshal(bytes, &wallet); err != nil {
		return responses.ServiceErrorToResponse(err)
	}

	if err := services.CreateWallet(wallet); err != nil {
		return responses.ServiceErrorToResponse(err)
	}

	wallet, err = services.GetWalletById(wallet.ID, []string{"User"})
	if err != nil {
		return responses.ServiceErrorToResponse(err)
	}

	return c.Status(fiber.StatusCreated).JSON(wallet)
}
