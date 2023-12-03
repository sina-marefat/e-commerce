package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Apply(server *fiber.App) {

	// Use Cors middleware
	server.Use(cors.New())
}
