package server

import (
	"fmt"

	"github.com/bagusyanuar/go-olin-bags-with-fiber/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

func Listen(cfg *config.Config, db *gorm.DB) {
	app := fiber.New(fiber.Config{
		TrustedProxies: []string{"127.0.0.1", "localhost"},
	})
	app.Use(recover.New())
	port := fmt.Sprintf(":%s", cfg.Port)
	app.Listen(port)
}
