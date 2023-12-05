package app

import (
	"e-commerce/config"
	v1 "e-commerce/internal/controller/http/v1"
	"e-commerce/internal/controller/http/v1/handlers"
	"e-commerce/pkg/fiber"
	"fmt"

	gofiber "github.com/gofiber/fiber/v2"
)

func StartHttpServer() {

	// Create Server
	server := fiber.NewServer(ServerConfig())

	// Register Routes
	RegisterRoutes(server.App)

	// Start Server
	fiber.StartServer(server)
}

// Register all domains routes
func RegisterRoutes(server *gofiber.App) {
	v1.RegisterRoutes(server, handlers.NewV1Handler(nil))
}

func ServerConfig() fiber.HTTPServerConfig {
	host := config.GetString("server.http.host")
	port := config.GetInt("server.http.port")
	fullAddress := fmt.Sprintf("%s:%d", host, port)

	return fiber.HTTPServerConfig{
		Host:        host,
		Port:        port,
		FullAddress: fullAddress,
	}
}
