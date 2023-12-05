package handlers

import "github.com/gofiber/fiber/v2"

type V1Handler struct {
}

func (h *V1Handler) Signup(c *fiber.Ctx) error {
	return nil
}

func (h *V1Handler) Login(c *fiber.Ctx) error {
	return nil
}

func NewV1Handler() *V1Handler {
	return &V1Handler{}
}
