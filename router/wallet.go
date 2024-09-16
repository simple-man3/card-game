package router

import (
	"card-game/controller"
	"card-game/middleware"
	"github.com/gofiber/fiber/v2"
)

func initWalletRouters(router fiber.Router) {
	walletController := controller.NewWalletController()

	group := router.Group("/wallet", middleware.CheckAuth)

	//router.Get("/:id<int>", controller.GetWallet)
	group.Post("/", walletController.CreateWallet)
}
