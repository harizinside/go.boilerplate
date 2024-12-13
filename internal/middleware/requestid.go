package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// RequestIDConfig returns Fiber Request ID middleware configuration.
func RequestIDConfig() fiber.Handler {
	return requestid.New()
}
