package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRodada(t *testing.T) {

	t.Run("Nova Rodada", func(t *testing.T) {
		rodada, err := NewRodada("Rodada 01", 10)
		assert.Nil(t, err)
		assert.NotNil(t, rodada)
		assert.NotEmpty(t, rodada.ID)
		assert.NotEmpty(t, rodada.Nome)
		assert.Equal(t, "Rodada 01", rodada.Nome)
		assert.Equal(t, 10, rodada.Peso)
		assert.NotEmpty(t, rodada.Criado)
		assert.Empty(t, rodada.Alterado)
		assert.Equal(t, true, rodada.Status)
	})

	t.Run("Alterando status", func(t *testing.T) {
		rodada, _ := NewRodada("Rodada 01", 10)
		assert.NotNil(t, rodada)

		rodada.Disable()
		assert.Equal(t, false, rodada.Status)

		rodada.Enable()
		assert.Equal(t, true, rodada.Status)
	})

}
