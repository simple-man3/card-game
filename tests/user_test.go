package tests

import (
	"bytes"
	"card-game/controller"
	"card-game/models"
	"card-game/requests"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"
)

type MockUserService struct {
	mock.Mock
}

func (mock *MockUserService) CreateUser(user *models.User) error {
	return mock.Called(user).Error(0)
}

func TestCreateUser(t *testing.T) {
	mockUserService := new(MockUserService)

	userController := controller.NewUserController()

	//if database.DBConn == nil {
	//	fmt.Println("error!!!!")
	//}

	app := fiber.New()
	app.Post("/api/v1/user", userController.CreateUser)

	request := requests.CreateUserRequest{
		Name:  "custom-name",
		Email: "example@example.com",
	}

	body, _ := json.Marshal(request)

	mockUserService.On("CreateUser", &request).Return(nil)

	req := httptest.NewRequest("POST", "/api/v1/user", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
}
