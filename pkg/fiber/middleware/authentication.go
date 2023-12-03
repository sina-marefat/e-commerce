package middleware

import (
	"e-commerce/config"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func JWTAuth() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.GetString("opts.jwt.privatekey"))},
	})
}
