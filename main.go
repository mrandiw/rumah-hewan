package main

import (
	"fmt"
	"log"
	"rumah-hewan/config"
	"rumah-hewan/router"
)

func main() {
	fmt.Println("Server is starting")

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("config tidak di temukan", err)
	}

	e := router.New()

	e.Start(config.App.Host)
}
