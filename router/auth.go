package router

import (
	"card-game/controller"
	"github.com/gofiber/fiber/v2"
)

func initAuthRouters(router fiber.Router) {
	authController := controller.NewAuthController()

	router.Post("/login", authController.Login)
}
