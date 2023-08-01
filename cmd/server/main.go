package main

import (
	"fmt"

	"github.com/bagusyanuar/go-olin-bags-with-fiber/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg, err := config.NewConfig(".env")
	if err != nil {
		panic(err)
	}
	app := fiber.New()
	port := fmt.Sprintf(":%s", cfg.Port)
	app.Listen(port)
}
