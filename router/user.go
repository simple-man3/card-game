package router

import (
	"card-game/controller"
	"github.com/gofiber/fiber/v2"
)

func initUserRouters(router fiber.Router) {
	router.Post("/user", controller.CreateUser)
	router.Patch("/user/:id<int>", controller.PatchUser)
}
