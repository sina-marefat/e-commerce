package app

import (
	"e-commerce/config"
	v1 "e-commerce/internal/controller/http/v1"
	"e-commerce/pkg/fiber"
	"e-commerce/pkg/fiber/middleware"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	gofiber "github.com/gofiber/fiber/v2"
)

// GracefullShutdownTimeout gracefull shutdown timeout
const GracefullShutdownTimeout = 10 * time.Second

func NewHttpServer() {

	address := fmt.Sprintf("%s:%d", config.GetString("service.http.host"), config.GetInt("server.http.port"))

	server := fiber.NewServer()

	middleware.Apply(server)

	v1.RegisterRoutes(server)

	go func() {
		err := server.Listen(address)
		if err != nil {
			log.Fatalf("Failed to start server : %s", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	if err := server.ShutdownWithTimeout(GracefullShutdownTimeout); err != nil {
		log.Fatal(err)
	}

}

func RegisterRoutes(server *gofiber.App) {
	v1.RegisterRoutes(server)
}
