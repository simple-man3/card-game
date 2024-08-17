package validator

import (
	"card-game/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validator validator.Validate

func InitInstanceValidator() {
	instance := validator.New()

	Validator = *instance
}

func InitValidator() error {
	InitInstanceValidator()
	SetErrorResponse()

	if err := RegisterCustomValidators(); err != nil {
		return err
	}

	return nil
}

func SetErrorResponse() {
	fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(responses.GlobalErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		},
	})
}

func RegisterCustomValidators() error {
	err := Validator.RegisterValidation("my-custom-valid", func(fl validator.FieldLevel) bool {
		return CheckUserExist(fl)
	})

	if err != nil {
		return err
	}

	return nil
}

func GetValidationErrResponse123(errs error) []responses.ErrorResponse {
	var errorResponses []responses.ErrorResponse
	for _, err := range errs.(validator.ValidationErrors) {
		var element responses.ErrorResponse
		element.FailedField = err.Field()
		element.Tag = err.Tag()
		element.Value = err.Value()
		errorResponses = append(errorResponses, element)
	}

	return errorResponses
}

func GetValidationErrResponse(errs error, c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).
		JSON(GetValidationErrResponse123(errs))
}
