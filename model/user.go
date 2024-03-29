package model

import (
	"time"

	"github.com/bagusyanuar/go-olin-bags-with-fiber/common"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const (
	UserTableName = "users"
)

type User struct {
	ID       uuid.UUID      `json:"id"`
	Email    string         `json:"email"`
	Username string         `json:"username"`
	Password *string        `json:"password"`
	Roles    datatypes.JSON `json:"roles"`
	common.WithTimestampsModel
	// Agent *Agent `gorm:"foreignKey:UserID" json:"agent"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return
}

func (User) TableName() string {
	return UserTableName
}
