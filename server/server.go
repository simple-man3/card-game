package server

import (
	"card-game/config"
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
		serverInstance = &Server{app: fiber.New()}
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
