package repositories

import (
	"github.com/bagusyanuar/go-olin-bags-with-fiber/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindByID(id string) (*model.User, error)
}

type User struct {
	Database *gorm.DB
}

// FindAll implements UserRepository.
func (r *User) FindAll() ([]model.User, error) {
	var data []model.User
	if err := r.Database.Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements UserRepository.
func (r *User) FindByID(id string) (*model.User, error) {
	data := new(model.User)
	if err := r.Database.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewUser(db *gorm.DB) UserRepository {
	return &User{
		Database: db,
	}
}
