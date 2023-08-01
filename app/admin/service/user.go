package service

import (
	"github.com/bagusyanuar/go-olin-bags-with-fiber/app/admin/repositories"
	"github.com/bagusyanuar/go-olin-bags-with-fiber/model"
)

type UserService interface {
	FindAll() ([]model.User, error)
	FindByID(id string) (*model.User, error)
}

type User struct {
	UserRepository repositories.UserRepository
}

// FindAll implements UserService.
func (svc *User) FindAll() ([]model.User, error) {
	return svc.UserRepository.FindAll()
}

// FindByID implements UserService.
func (svc *User) FindByID(id string) (*model.User, error) {
	return svc.UserRepository.FindByID(id)
}

func NewUser(userRepository repositories.UserRepository) UserService {
	return &User{
		UserRepository: userRepository,
	}
}
