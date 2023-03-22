package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {

	t.Run("Nova usuário", func(t *testing.T) {
		user, err := NewUser("Bruno", "brunolimame@gmail.com", "123456")
		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.NotEmpty(t, user.ID)
		assert.NotEmpty(t, user.Nome)
		assert.NotEmpty(t, user.Email)
		assert.NotEmpty(t, user.Password)
		assert.Equal(t, "Bruno", user.Nome)
		assert.Equal(t, "brunolimame@gmail.com", user.Email)
	})

	t.Run("Validar Password", func(t *testing.T) {
		user, err := NewUser("Bruno", "brunolimame@gmail.com", "123456")
		assert.Nil(t, err)
		assert.True(t, user.ValidarSenha("123456"))
		assert.False(t, user.ValidarSenha("1234567"))
		assert.NotEqual(t, "123456", user.Password)
	})
}
