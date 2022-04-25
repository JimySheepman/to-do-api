package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JimySheepman/to-do-api/internal/api/router"
	"github.com/JimySheepman/to-do-api/internal/infrastructure/persistence"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const idleTimeout = 10 * time.Second

func GracefulShutdown(app *fiber.App, port string) {
	// Listen from a different goroutine
	go func() {
		if err := app.Listen(port); err != nil {
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

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})

	app.Use(cors.New())

	persistence.ConnectDB()

	router.SetupRoutes(app)

	GracefulShutdown(app, ":8080")
}
