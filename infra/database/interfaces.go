package database

import "bolao/internal/entity"

type UserInterface interface {
	Create(user *entity.UserEntity) error
	FindByEmail(email string) (*entity.UserEntity, error)
}
