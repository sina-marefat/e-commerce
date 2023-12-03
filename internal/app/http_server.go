package app

import (
	v1 "e-commerce/internal/controller/http/v1"
	"e-commerce/pkg/fiber"
	"log"
)

func NewHttpServer() {

	server := fiber.NewServer()

	v1.RegisterRoutes(server)

	err := server.Listen(":3000")
	if err != nil {
		log.Fatalf("Failed to start server : %s", err)
	}
}
