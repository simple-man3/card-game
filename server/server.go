package server

import (
	"card-game/config"
	"card-game/responses"
	"card-game/router"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"sync"
)

type Server struct {
	app *fiber.App
}

var (
	serverInstance *Server
	once           sync.Once
)

func NewServer() *Server {
	once.Do(func() {
		/**
		 * toDo
		 *  [] необходимо переработать глобальный обработчик ошибок
		 */
		serverInstance = &Server{
			app: fiber.New(fiber.Config{
				ErrorHandler: func(c *fiber.Ctx, err error) error {
					return c.Status(fiber.StatusBadRequest).JSON(responses.GlobalErrorHandlerResp{
						Success: false,
						Message: err.Error(),
					})
				},
			}),
		}
	})

	return serverInstance
}

func (s *Server) InitRouters() error {
	router.InitRouters(s.app)
	return nil
}

func (s *Server) Start() error {
	env, err := config.GetInstanceEnv()
	if err != nil {
		return err
	}

	port := fmt.Sprintf(":%s", env.AppPort)
	if err := s.app.Listen(port); err != nil {
		return err
	}

	return nil
}
