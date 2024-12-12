package api

import (
	"go.boilerplate/internal/app/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(app *fiber.App, db *mongo.Database, rdb *redis.Client) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	iniRoutes := app.Group("/auth")
	auth.AuthRoutes(iniRoutes, db)
}
