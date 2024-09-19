package router

import (
	"card-game/controller"
	"card-game/middleware"
	"github.com/gofiber/fiber/v2"
)

func initWalletRouters(router fiber.Router) {
	walletController := controller.NewWalletController()

	group := router.Group("/wallet", middleware.CheckAuth)

	group.Get("/:id<int>", walletController.Get)
	group.Post("/", walletController.CreateWallet)
	group.Post("/put-money", walletController.PutMoney)
}
