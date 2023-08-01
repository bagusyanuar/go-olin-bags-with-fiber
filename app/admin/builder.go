package admin

import (
	"github.com/bagusyanuar/go-olin-bags-with-fiber/app/admin/handlers"
	"github.com/bagusyanuar/go-olin-bags-with-fiber/app/admin/repositories"
	"github.com/bagusyanuar/go-olin-bags-with-fiber/app/admin/service"
	"github.com/bagusyanuar/go-olin-bags-with-fiber/common"
	"github.com/bagusyanuar/go-olin-bags-with-fiber/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Builder struct {
	Database *gorm.DB
	Config   *config.Config
	Router   fiber.Router
}

func NewBuilder(db *gorm.DB, cfg *config.Config, router fiber.Router) Builder {
	return Builder{Database: db, Config: cfg, Router: router}
}

func (b *Builder) Build() {
	userRepository := repositories.NewUser(b.Database)

	userService := service.NewUser(userRepository)

	userHandler := handlers.NewUser(userService, b.Router)

	handlers := []any{
		&userHandler,
	}

	common.RegisterRoutes(handlers...)
}
