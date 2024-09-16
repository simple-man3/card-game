package responses

import "github.com/gofiber/fiber/v2"

type (
	LoginResponse struct {
		Token string `json:"token"`
	}
)

func ForbiddenResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"message": "Not auth",
	})
}
