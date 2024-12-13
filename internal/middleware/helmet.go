package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

// HelmetConfig returns Fiber Helmet middleware configuration.
func HelmetConfig() fiber.Handler {
	return helmet.New()
}
