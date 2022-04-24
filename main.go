package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

const idleTimeout = 10 * time.Second

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	// Listen from a different goroutine
	go func() {
		if err := app.Listen(":8080"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")
	fmt.Println("Fiber was successful shutdown.")
}
