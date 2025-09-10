package config

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

// FiberConfig returns fiber.Config with environment-based settings
func FiberConfig() *fiber.Config {
	// Common config
	cfg := fiber.Config{
		AppName:      "Darinol",
		BodyLimit:    4 * 1024 * 1024, // 4 MB
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	return &cfg
}
