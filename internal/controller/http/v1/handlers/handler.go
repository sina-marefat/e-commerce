package handlers

import (
	"e-commerce/internal/controller/http/v1/dto/request"
	"e-commerce/internal/usecase"
	"e-commerce/pkg/fiber"
	"e-commerce/pkg/validator"

	gofiber "github.com/gofiber/fiber/v2"
)

type V1Handler struct {
	authUsecase usecase.AuthUsecase
}

func (h *V1Handler) Signup(c *gofiber.Ctx) error {
	signUpRequest := new(request.CreateUser)

	if err := c.BodyParser(signUpRequest); err != nil {
		return err
	}

	if errs := validator.Validate(signUpRequest); len(errs) > 0 && errs[0].Error {
		return fiber.StatusBadRequest(c, map[string]interface{}{
			"message": "validation failed",
			"fields":  errs,
		})
	}

	err := h.authUsecase.SignUp(c.Context(), *signUpRequest)

	if err != nil {
		return err
	}

	return nil
}

func (h *V1Handler) Login(c *gofiber.Ctx) error {
	loginRequest := new(request.Login)

	if err := c.BodyParser(loginRequest); err != nil {
		return err
	}

	if errs := validator.Validate(loginRequest); len(errs) > 0 && errs[0].Error {
		return fiber.StatusBadRequest(c, map[string]interface{}{
			"message": "validation failed",
			"fields":  errs,
		})
	}

	_, err := h.authUsecase.Login(c.Context(), *loginRequest)

	if err != nil {
		return err
	}

	return nil
}

func NewV1Handler(authUsecase usecase.AuthUsecase) *V1Handler {
	return &V1Handler{}
}
