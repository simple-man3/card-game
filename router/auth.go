package router

import (
	"card-game/controller"
	"github.com/gofiber/fiber/v2"
)

func initAuthRouters(router fiber.Router) {
	router.Post("/login", controller.Login)
}
