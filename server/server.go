package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/bytedance/sonic"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	HelloworldAPI "github.com/kuzuru/KIS/api/helloworld"
)

func Run(port string) {
	// Decalring app router
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		JSONEncoder:   sonic.Marshal,
		JSONDecoder:   sonic.Unmarshal,
	})

	// Using middlewares
	app.Use(
		recover.New(),
	)

	// Defining /api group
	api := app.Group("/api")

	// Register HTTP handlers
	HelloworldAPI.RegisterHTTPEndpoint(api)

	// Running server in background
	go func() {
		if err := app.Listen(":" + port); err != nil {
			log.Panicf("[error] app.Listen: %s", err)
		}
	}()

	// Waiting for quit signal on exit
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	_, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
}
