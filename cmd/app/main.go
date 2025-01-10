package main

import (
	"log"
	"os"

	"go.boilerplate/api"
	"go.boilerplate/internal/config"
	"go.boilerplate/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Mongodb database connection
	err = config.ConnectMongo(os.Getenv("DB_HOST"), os.Getenv("DB_DATABASE"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer config.DisconnectMongo()

	db := config.GetDatabase()
	log.Printf("Using database: %s", db.Name())

	// Redis cache connection
	if err := config.ConnectRedis(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PASSWORD")); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer config.DisconnectRedis()

	rdb := config.GetRedisClient()
	log.Printf("Redis connected")

	app := config.FiberConfig()

	app.Use(middleware.RecoverConfig())

	app.Use(middleware.CacheConfig())

	app.Use(middleware.RequestIDConfig())

	app.Use(middleware.LoggerConfig())

	app.Use(middleware.LimiterConfig())

	app.Use(middleware.CompressConfig())

	app.Use(middleware.CorsConfig())

	app.Use(middleware.HelmetConfig())

	app.Use(middleware.ETagConfig())

	app.Use(middleware.FileSystemConfig())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome Golang on Fiber")
	})

	api.SetupRoutes(app, db, rdb)

	if err := app.Listen(os.Getenv("API_URI")); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
