package router

import (
	"card-game/controller"
	"github.com/gofiber/fiber/v2"
)

func initWalletRouters(router fiber.Router) {
	walletController := controller.NewWalletController()

	//router.Get("/wallet/:id<int>", controller.GetWallet)
	router.Post("/wallet", walletController.CreateWallet)
}
