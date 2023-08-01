package handlers

import (
	"github.com/bagusyanuar/go-olin-bags-with-fiber/app/admin/service"
	"github.com/bagusyanuar/go-olin-bags-with-fiber/common"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type User struct {
	UserService service.UserService
	Router      fiber.Router
}

func NewUser(userService service.UserService, router fiber.Router) User {
	return User{UserService: userService, Router: router}
}

func (u *User) Routes() {
	route := u.Router.Group("/user")
	route.Get("/", u.FindAll)
	route.Get("/:id", u.FindByID)
}

func (u *User) FindAll(ctx *fiber.Ctx) error {
	data, err := u.UserService.FindAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.APIResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})

	}
	return ctx.Status(fiber.StatusOK).JSON(common.APIResponse{
		Code:    fiber.StatusOK,
		Message: "success",
		Data:    data,
	})
}

func (u *User) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data, err := u.UserService.FindByID(id)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return ctx.Status(fiber.StatusNotFound).JSON(common.APIResponse{
				Code:    fiber.StatusNotFound,
				Message: "data not found",
				Data:    nil,
			})

		default:
			return ctx.Status(fiber.StatusInternalServerError).JSON(common.APIResponse{
				Code:    fiber.StatusInternalServerError,
				Message: err.Error(),
				Data:    nil,
			})
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(common.APIResponse{
		Code:    fiber.StatusOK,
		Message: err.Error(),
		Data:    data,
	})
}
