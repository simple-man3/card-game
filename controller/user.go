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

type UserController struct {
	UserService *services.UserService
}

func NewUserController() *UserController {
	userService := services.NewUserService()

	return &UserController{
		UserService: userService,
	}
}

func (controller *UserController) CreateUser(c *fiber.Ctx) error {
	var request requests.CreateUserRequest

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

	var user models.User
	if err := json.Unmarshal(bytes, &user); err != nil {
		return responses.ServiceErrorToResponse(err)
	}

	if err := controller.UserService.CreateUser(&user); err != nil {
		return responses.ServiceErrorToResponse(err)
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (controller *UserController) PatchUser(c *fiber.Ctx) error {
	var request requests.PatchUserRequest
	var id, _ = c.ParamsInt("id")

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

	var user models.User
	if err := json.Unmarshal(bytes, &user); err != nil {
		return responses.ServiceErrorToResponse(err)
	}

	if err := controller.UserService.UpdateUser(&user, uint(id)); err != nil {
		return responses.ServiceErrorToResponse(err)
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (controller *UserController) GetUser(c *fiber.Ctx) error {
	var id, _ = c.ParamsInt("id")

	user, err := controller.UserService.GetUserById(uint(id))
	if err != nil {
		return responses.ServiceErrorToResponse(err)
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (controller *UserController) DeleteUser(c *fiber.Ctx) error {
	var id, _ = c.ParamsInt("id")
	user := models.User{ID: uint(id)}

	err := controller.UserService.DeleteUser(user)
	if err != nil {
		return responses.ServiceErrorToResponse(err)
	}

	return c.SendStatus(fiber.StatusOK)
}
