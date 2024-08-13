package main

import (
	"fmt"
	"github.com/tegarsubkhan236/go-rbac/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	viperConfig := config.NewViper()
	db := config.NewDatabase(viperConfig)
	redis := config.NewRedis(viperConfig)
	app := config.NewFiber(viperConfig)
	validator := config.NewValidator(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		Config:    viperConfig,
		DB:        db,
		Redis:     redis,
		App:       app,
		Validator: validator,
	})

	go func() {
		webPort := viperConfig.GetInt("web.port")
		err := app.Listen(fmt.Sprintf(":%d", webPort))
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	_ = <-c

	fmt.Println("\nGracefully shutting down...")
	_ = app.Shutdown()
	fmt.Println("Fiber was successful shutdown.")
}
