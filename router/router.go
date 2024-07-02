package router

import (
	"card-game/controller"
	"github.com/gofiber/fiber/v2"
)

func InitRouters(app *fiber.App) {
	app.Get("/health", controller.Health)
}
