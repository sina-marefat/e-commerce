package fiber

import "github.com/gofiber/fiber/v2"

type CustomResponse struct {
	StatusCode int         `json:"status"`
	Data       interface{} `json:"data,omitempty"`
	Message    string      `json:"message"`
}

func StatusOk(c *fiber.Ctx, data interface{}, customStatus ...int) error {
	// Set Content-Type: text/plain; charset=utf-8
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	code := fiber.StatusOK

	if len(customStatus) > 0 {
		code = customStatus[0]
	}

	response := CustomResponse{
		StatusCode: code,
		Message:    "OK",
		Data:       data,
	}

	// Return status code with error message
	return c.Status(fiber.StatusOK).JSON(response)
}

func StatusServerError(c *fiber.Ctx, data interface{}, customStatus ...int) error {
	// Set Content-Type: text/plain; charset=utf-8
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	code := fiber.StatusInternalServerError

	if len(customStatus) > 0 {
		code = customStatus[0]
	}

	response := CustomResponse{
		StatusCode: code,
		Message:    "Internal Server Error",
		Data:       data,
	}

	// Return status code with error message
	return c.Status(fiber.StatusInternalServerError).JSON(response)
}

func StatusBadRequest(c *fiber.Ctx, data interface{}, customStatus ...int) error {
	// Set Content-Type: text/plain; charset=utf-8
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	code := fiber.StatusBadRequest

	if len(customStatus) > 0 {
		code = customStatus[0]
	}

	response := CustomResponse{
		StatusCode: code,
		Message:    "Bad Request",
		Data:       data,
	}

	// Return status code with error message
	return c.Status(fiber.StatusBadRequest).JSON(response)
}
