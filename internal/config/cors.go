package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CorsConfig() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "*",                            // Allow all origins (you can restrict this to specific domains)
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",  // Allowed HTTP methods
		AllowHeaders: "Origin, Content-Type, Accept", // Allowed request headers
	})
}
