package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/tegarsubkhan236/redis-jwt-auth/internal/pkg/config"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.NewRedis()
	defer func() {
		if err := config.CloseRedis(); err != nil {
			log.Printf("Error closing Redis: %v", err)
		}
	}()

	app := fiber.New(fiber.Config{
		IdleTimeout: 5 * time.Second,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	_ = <-c

	fmt.Println("\nGracefully shutting down...")
	_ = app.Shutdown()
	fmt.Println("Fiber was successful shutdown.")
}
