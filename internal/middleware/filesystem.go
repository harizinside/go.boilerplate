package middleware

import (
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

// FilesystemMiddleware serves static files from the "public" directory.
func FileSystemConfig() fiber.Handler {
	// Ensure the "public" directory exists
	if _, err := os.Stat("./public"); os.IsNotExist(err) {
		if err := os.Mkdir("./public", os.ModePerm); err != nil {
			panic("Failed to create public directory: " + err.Error())
		}
	}

	return filesystem.New(filesystem.Config{
		Root:   http.Dir("./public"), // Serve files from "./public" directory
		Browse: true,                 // Enable directory browsing
		Index:  "index.html",         // Serve "index.html" for directory access
		MaxAge: 3600,                 // Cache control max age
	})
}
