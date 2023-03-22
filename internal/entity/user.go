package entity

import (
	"bolao/pkg/entity"

	"golang.org/x/crypto/bcrypt"
)

type UserEntity struct {
	ID       entity.ID `json:"id"`
	Nome     string    `json:"nome"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	Criado   string    `json:"criado"`
	Alterado string    `json:"alterado"`
}

func NewUser(nome, email, password string) (*UserEntity, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &UserEntity{
		ID:       entity.NewID(),
		Nome:     nome,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *UserEntity) ValidarSenha(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
