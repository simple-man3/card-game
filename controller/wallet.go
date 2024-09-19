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

type WalletController struct {
	walletService *services.WalletService
}

func NewWalletController() *WalletController {
	walletService := services.NewWalletService()

	return &WalletController{
		walletService: walletService,
	}
}

func (wc WalletController) CreateWallet(c *fiber.Ctx) error {
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

	if err := wc.walletService.CreateWallet(wallet); err != nil {
		return responses.ServiceErrorToResponse(err)
	}

	wallet, err = wc.walletService.GetWalletById(wallet.ID, models.AuthUser.ID, []string{"User"})
	if err != nil {
		return responses.ServiceErrorToResponse(err)
	}

	return c.Status(fiber.StatusCreated).JSON(wallet)
}

func (wc WalletController) Get(c *fiber.Ctx) error {
	var id, _ = c.ParamsInt("id")

	wallet, err := wc.walletService.GetWalletById(uint(id), models.AuthUser.ID, []string{"User"})
	if err != nil {
		return responses.ServiceErrorToResponse(err)
	}

	return c.Status(fiber.StatusOK).JSON(wallet)
}

func (wc WalletController) PutMoney(c *fiber.Ctx) error {
	var request requests.PutMoneyRequest

	if err := c.BodyParser(&request); err != nil {
		return responses.BodyParseErrToResponse()
	}

	if errs := validator.Validator.Struct(request); errs != nil {
		return responses.ValidationErrToResponse(errs, c)
	}

	if err := wc.walletService.PutMoney(request.Amount, *models.AuthUser); err != nil {
		return responses.ValidationErrToResponse(err, c)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (wc WalletController) WithdrawMoney(c *fiber.Ctx) error {
	var request requests.WithDrawMoneyRequest

	if err := c.BodyParser(&request); err != nil {
		return responses.BodyParseErrToResponse()
	}

	if errs := validator.Validator.Struct(request); errs != nil {
		return responses.ValidationErrToResponse(errs, c)
	}

	if err := wc.walletService.WithdrawMoney(request.Amount, *models.AuthUser); err != nil {
		return responses.ServiceErrorToResponse(err)
	}

	return c.SendStatus(fiber.StatusOK)
}
