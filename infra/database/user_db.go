package database

import (
	"bolao/internal/entity"

	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(user *entity.UserEntity) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByEmail(email string) (*entity.UserEntity, error) {

}
