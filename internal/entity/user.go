package entity

import (
	"bolao/pkg/entity"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"`
	Nome     string    `json:"nome"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUser(nome, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       entity.NewID(),
		Nome:     nome,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *User) ValidarSenha(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
