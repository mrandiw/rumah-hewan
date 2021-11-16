package main

import (
	"fmt"
	"rumah-hewan/router"
)

func main() {
	fmt.Println("Server is starting")

	e := router.New()

	e.Start(":8080")
}
