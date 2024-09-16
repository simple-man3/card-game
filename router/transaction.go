package router

import (
	"card-game/controller"
	"card-game/middleware"
	"github.com/gofiber/fiber/v2"
)

func initTransactionRouters(router fiber.Router) {
	transactionController := controller.NewTransactionController()

	group := router.Group("/transaction", middleware.CheckAuth)

	group.Post("/put-money", transactionController.PutMoney)
}
