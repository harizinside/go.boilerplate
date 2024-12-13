package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// LimiterConfig returns Fiber Rate Limiter middleware configuration.
func LimiterConfig() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        10,               // Maximum number of requests
		Expiration: 30 * time.Second, // Time window for the request limit
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests. Please try again later.",
			})
		},
	})
}
