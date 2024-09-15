package main

import (
	"card-game/config"
	"card-game/database"
	_ "card-game/docs"
	"card-game/server"
	"card-game/session"
	"card-game/validator"
	"log"
)

// @title Card Game API
// @version 1.0
// @host localhost:8080
// @BasePath /api
func main() {
	serv := server.NewServer()

	if err := serv.InitRouters(); err != nil {
		log.Fatal(err)
	}

	if err := config.InitEnv(); err != nil {
		log.Fatal(err)
	}

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	if err := database.AutoMigrate(); err != nil {
		log.Fatal(err)
	}

	if err := validator.InitValidator(); err != nil {
		log.Fatal(err)
	}

	session.InitSession()

	if err := serv.Start(); err != nil {
		log.Fatal(err)
	}
}
