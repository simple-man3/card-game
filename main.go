package main

import (
	"card-game/config"
	"card-game/server"
	"log"
)

func main() {
	serv := server.NewServer()

	if err := serv.InitRouters(); err != nil {
		log.Fatal(err)
	}

	if _, err := config.NewEnv(); err != nil {
		log.Fatal(err)
	}

	if err := serv.Start(); err != nil {
		log.Fatal(err)
	}
}
