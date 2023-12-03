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

type HTTPServerConfig struct {
	Host        string
	Port        int
	FullAddress string
}

// GracefullShutdownTimeout gracefull shutdown timeout
const GracefullShutdownTimeout = 10 * time.Second

func NewHttpServer() {

	serverConfig := ServerConfig()
	// Create Server
	server := fiber.NewServer()

	// Apply middlewares
	middleware.Apply(server)

	// Register Routes
	v1.RegisterRoutes(server)

	// Serve on goroutine to avoid blocking (for graceful shut down)
	go func() {
		err := server.Listen(serverConfig.FullAddress)
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

func ServerConfig() HTTPServerConfig {
	host := config.GetString("server.http.host")
	port := config.GetInt("server.http.port")
	fullAddress := fmt.Sprintf("%s:%d", host, port)

	return HTTPServerConfig{
		Host:        host,
		Port:        port,
		FullAddress: fullAddress,
	}
}
