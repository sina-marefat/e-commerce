package middleware

import (
	"e-commerce/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Apply(server *fiber.App) {

	// Use Cors middleware
	server.Use(cors.New())
	// How to change log level
	if utils.IsDebugMode() {
		server.Use(logger.New())
	}

}
