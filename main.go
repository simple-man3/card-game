package main

import (
	"card-game/config"
	"card-game/database"
	"card-game/server"
	"log"
)

func main() {
	serv := server.NewServer()

	if err := serv.InitRouters(); err != nil {
		log.Fatal(err)
	}

	if _, err := config.GetInstanceEnv(); err != nil {
		log.Fatal(err)
	}

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	if err := database.AutoMigrate(); err != nil {
		log.Fatal(err)
	}

	if err := serv.Start(); err != nil {
		log.Fatal(err)
	}
}
