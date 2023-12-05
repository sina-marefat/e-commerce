package v1

import (
	"e-commerce/internal/controller/http/v1/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(server *fiber.App, handler *handlers.V1Handler) {
	v1 := server.Group("/v1")
	{
		user := v1.Group("/user")
		{
			auth := user.Group("/auth")
			{
				auth.Post("/signup", handler.Signup)
				auth.Post("/login", handler.Login)
			}
		}
	}
}
