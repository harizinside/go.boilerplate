package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
)

// ETagMiddleware sets up the ETag middleware with default settings.
func ETagConfig() fiber.Handler {
	return etag.New()
}
