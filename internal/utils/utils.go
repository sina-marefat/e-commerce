package utils

import (
	"e-commerce/config"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func GetUserIDFromJWT(c *fiber.Ctx) (uint, error) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	uidStr := claims["uid"].(string)
	uid, err := strconv.ParseUint(uidStr, 10, 32)

	return uint(uid), err
}

func IsDebugMode() bool {
	return config.GetBool("debug_mode")
}
