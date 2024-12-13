package config

import (
	"time"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

// FiberConfig returns a configured Fiber application instance.
func FiberConfig() *fiber.App {
	return fiber.New(fiber.Config{
		JSONEncoder:   sonic.Marshal,        // Use Sonic for faster JSON encoding
		JSONDecoder:   sonic.Unmarshal,      // Use Sonic for faster JSON decoding
		ServerHeader:  "GolangFiberAPI/1.0", // Custom Server Header
		AppName:       "GolangFiberAPI/1.0", // Application Name
		CaseSensitive: true,                 // Enable case-sensitive routing
		StrictRouting: true,                 // Enable strict routing (e.g., /foo != /foo/)
		ReadTimeout:   10 * time.Second,     // Set request read timeout
		WriteTimeout:  10 * time.Second,     // Set response write timeout
		IdleTimeout:   30 * time.Second,     // Set idle timeout
		BodyLimit:     10 * 1024 * 1024,     // Set body size limit to 10 MB
		// Prefork:       true,              // Uncomment for multicore usage
	})
}
