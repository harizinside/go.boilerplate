package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

// CacheMiddleware sets up the cache middleware with default settings.
func CacheConfig() fiber.Handler {
	return cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			// Skip cache for specific requests (e.g., if requested by query parameter)
			return c.Query("skipCache") == "true"
		},
		Expiration:   10 * time.Minute, // Cache expiration time
		CacheControl: true,             // Enable Cache-Control header
	})
}
