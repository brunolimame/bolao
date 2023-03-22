package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {

	t.Run("Novo time", func(t *testing.T) {
		time, err := NewTime("Brasil", "escudo.jpg")
		assert.Nil(t, err)
		assert.NotNil(t, time)
		assert.NotEmpty(t, time.ID)
		assert.NotEmpty(t, time.Nome)
		assert.Equal(t, "Brasil", time.Nome)
		assert.Equal(t, "escudo.jpg", time.Escudo)
		assert.NotEmpty(t, time.Criado)
		assert.Empty(t, time.Alterado)
		assert.Equal(t, true, time.Status)
	})

	t.Run("Alterando status", func(t *testing.T) {
		time, _ := NewTime("Brasil", "escudo.jpg")
		assert.NotNil(t, time)

		time.Disable()
		assert.Equal(t, false, time.Status)

		time.Enable()
		assert.Equal(t, true, time.Status)
	})
}
