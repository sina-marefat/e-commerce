package v1

import "github.com/gofiber/fiber/v2"

type v1Handler struct {
}

func (h *v1Handler) signup(c *fiber.Ctx) error {
	return nil
}

func (h *v1Handler) login(c *fiber.Ctx) error {
	return nil
}

func RegisterRoutes(server *fiber.App) {
	handler := v1Handler{}
	v1 := server.Group("/v1")
	{
		user := v1.Group("/user")
		{
			auth := user.Group("/auth")
			{
				auth.Post("/signup", handler.signup)
				auth.Post("/login", handler.login)
			}
		}
	}
}
