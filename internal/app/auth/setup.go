package auth

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthRoutes(app fiber.Router, db *mongo.Database) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	app.Post("/sign-up", handler.SignUp)
	app.Post("/sign-in", handler.SignIn)
	app.Post("/reset-password", handler.Recovery)
	app.Post("/reset-password/:id", handler.ResetPassword)
}
