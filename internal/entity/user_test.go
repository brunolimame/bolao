package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {

	t.Run("Novo usuário", func(t *testing.T) {
		user, err := NewUser("Bruno", "brunolimame@gmail.com", "123456")
		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.NotEmpty(t, user.ID)
		assert.NotEmpty(t, user.Nome)
		assert.NotEmpty(t, user.Email)
		assert.NotEmpty(t, user.Password)
		assert.NotZero(t, len(user.Password))
		assert.NotEmpty(t, user.Salt)
		assert.Equal(t, "Bruno", user.Nome)
		assert.Equal(t, "brunolimame@gmail.com", user.Email)
		assert.NotEmpty(t, user.Criado)
		assert.Empty(t, user.Alterado)
		assert.Equal(t, true, user.Status)
	})

	t.Run("Validar Password", func(t *testing.T) {
		user, err := NewUser("Bruno", "brunolimame@gmail.com", "123456")
		assert.Nil(t, err)
		assert.False(t, user.ValidPassword("1234567"))
		assert.True(t, user.ValidPassword("123456"))
	})

	t.Run("Criando usuário com nome em branco", func(t *testing.T) {
		user, err := NewUser("", "brunolimame@gmail.com", "123456")
		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualError(t, err, UserEntityMsgErrorNomeRequerido)
	})

	t.Run("Criando usuário com e-mail em branco", func(t *testing.T) {
		user, err := NewUser("Bruno", "", "123456")
		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualError(t, err, UserEntityMsgErrorEmailRequerido)
	})

	t.Run("Criando usuário com senha em branco", func(t *testing.T) {
		user, err := NewUser("Bruno", "brunolimame@gmail.com", "")
		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualError(t, err, UserEntityMsgErrorPasswordRequerido)
	})
}
