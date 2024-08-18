package responses

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	ErrorResponse struct {
		FailedField string `json:"field"`
		Tag         string `json:"validation"`
		Value       any    `json:"value"`
	}
)

func generateErrorResponses(errs error) []ErrorResponse {
	var errorResponses []ErrorResponse
	for _, err := range errs.(validator.ValidationErrors) {
		var element ErrorResponse
		element.FailedField = err.Field()
		element.Tag = err.Tag()
		element.Value = err.Value()
		errorResponses = append(errorResponses, element)
	}

	return errorResponses
}

func BodyParseErrToResponse() error {
	return &fiber.Error{
		Code:    fiber.StatusBadRequest,
		Message: "Invalid request body",
	}
}

func ValidationErrToResponse(errs error, c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).
		JSON(generateErrorResponses(errs))
}

func ServiceErrorToResponse(err error) error {
	return &fiber.Error{
		Code:    fiber.StatusInternalServerError,
		Message: err.Error(),
	}
}
