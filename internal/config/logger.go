package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// LoggerConfig returns Fiber Logger middleware configuration.
func LoggerConfig() fiber.Handler {
	return logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n", // Log format
	})
}
