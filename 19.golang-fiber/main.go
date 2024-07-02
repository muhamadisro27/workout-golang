package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)


func main() {
	newApp := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		Prefork:      true,
	})

	if fiber.IsChild() {
		fmt.Println("I'm child process")
	} else {
		fmt.Println("I'm parent process")
	}

	err := newApp.Listen("localhost:4000")

	if err != nil {
		panic(err)
	}
}
