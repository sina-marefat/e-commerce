package v1

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(server *fiber.App) {
	server.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})
}
