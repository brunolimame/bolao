package entity

import (
	"bolao/pkg/entity"
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const userEntityDefaultCostPassword = bcrypt.DefaultCost
const UserEntityMsgErrorNomeRequerido = "Nome do usuário não definido"
const UserEntityMsgErrorEmailRequerido = "E-mail não definido"
const UserEntityMsgErrorPasswordRequerido = "Senha não definida"

type UserEntity struct {
	ID       entity.ID `json:"id"`
	Nome     string    `json:"nome"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	Salt     string    `json:"-"`
	Criado   time.Time `json:"criado"`
	Alterado time.Time `json:"alterado"`
	Status   bool      `json:"status"`
}

func NewUser(nome, email, password string) (*UserEntity, error) {

	newSalt, err := genSalt()
	if err != nil {
		return nil, err
	}

	newUser := &UserEntity{
		ID:       entity.NewID(),
		Nome:     nome,
		Email:    email,
		Password: password,
		Salt:     newSalt,
		Criado:   time.Now(),
		Alterado: time.Time{},
		Status:   true,
	}

	err = newUser.Validate()

	if err != nil {
		return nil, err
	}

	passwordCrypt, err := newUser.EncryptPassword(newUser.Password + newUser.Salt)

	if err != nil {
		return nil, err
	}
	newUser.Password = string(passwordCrypt)

	return newUser, nil
}

func (u *UserEntity) Validate() error {
	if len(u.Nome) <= 0 {
		return errors.New(UserEntityMsgErrorNomeRequerido)
	}
	if len(u.Email) <= 0 {
		return errors.New(UserEntityMsgErrorEmailRequerido)
	}
	if len(u.Password) <= 0 {
		return errors.New(UserEntityMsgErrorPasswordRequerido)
	}
	return nil
}

func genSalt() (string, error) {
	salt := make([]byte, userEntityDefaultCostPassword)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", salt), nil
}

func (u *UserEntity) EncryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), userEntityDefaultCostPassword)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (u *UserEntity) ValidPassword(password string) bool {
	testPassword := password + u.Salt
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(testPassword))
	return err == nil
}
