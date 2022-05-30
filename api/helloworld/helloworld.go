package helloworld

import "github.com/gofiber/fiber/v2"

func RegisterHTTPEndpoint(api fiber.Router) {
	api.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
}
