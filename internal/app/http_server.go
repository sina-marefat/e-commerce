package app

import (
	"e-commerce/config"
	v1 "e-commerce/internal/controller/http/v1"
	"e-commerce/pkg/fiber"
	"fmt"
	"log"
)

func NewHttpServer() {

	address := fmt.Sprintf("%s:%d", config.GetString("service.http.host"), config.GetInt("server.http.port"))

	server := fiber.NewServer()

	v1.RegisterRoutes(server)

	err := server.Listen(address)

	if err != nil {
		log.Fatalf("Failed to start server : %s", err)
	}
}
