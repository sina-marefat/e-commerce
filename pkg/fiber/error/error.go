package error

import (
	"e-commerce/pkg/fiber"
	"errors"

	gofiber "github.com/gofiber/fiber/v2"
)

func CustomErrorHandler(c *gofiber.Ctx, err error) error {
	// Status code defaults to 500
	code := gofiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *gofiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	// Set Content-Type: text/plain; charset=utf-8
	c.Set(gofiber.HeaderContentType, gofiber.MIMEApplicationJSONCharsetUTF8)

	response := fiber.CustomResponse{
		StatusCode: code,
		Message:    "Internal Server Error",
		Data: map[string]interface{}{
			"message": err.Error(),
		},
	}

	// Return status code with error message
	return c.Status(code).JSON(response)

}
