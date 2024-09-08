package router

import (
	"card-game/controller"
	"github.com/gofiber/fiber/v2"
)

func initUserRouters(router fiber.Router) {
	userController := controller.NewUserController()

	router.Post("/user", userController.CreateUser)
	router.Patch("/user/:id<int>", userController.PatchUser)
	router.Get("/user/:id<int>", userController.GetUser)
	router.Delete("/user/:id<int>", userController.DeleteUser)
}
