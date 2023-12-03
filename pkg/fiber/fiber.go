package fiber

import (
	"e-commerce/pkg/fiber/middleware"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
)

const GracefullShutdownTimeout = 10 * time.Second

type FiberServer struct {
	*fiber.App
	config HTTPServerConfig
}

type HTTPServerConfig struct {
	Host        string
	Port        int
	FullAddress string
}

func NewServer(config HTTPServerConfig) *FiberServer {

	// Create Fiber instance
	server := fiber.New()

	// Apply middlewares
	middleware.Apply(server)

	return &FiberServer{server, config}
}

func StartServer(server *FiberServer) {

	// Serve on goroutine to avoid blocking (for graceful shut down)
	go func() {
		err := server.Listen(server.config.FullAddress)
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
