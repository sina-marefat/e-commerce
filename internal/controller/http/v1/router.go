package v1

import "github.com/gofiber/fiber/v2"

type v1Handler struct {
}

func RegisterRoutes(server *fiber.App) {
	// handler := v1Handler{}
	server.Group("/v1")
}
