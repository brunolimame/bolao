package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalpite(t *testing.T) {

	t.Run("Criando novo palpite", func(t *testing.T) {
		palpite, err := NewPalpite("p1", "j1", 1, 0)
		assert.Nil(t, err)
		assert.NotNil(t, palpite)
		assert.NotEmpty(t, palpite.ID)
		assert.NotEmpty(t, palpite.PlayerID)
		assert.Equal(t, "p1", palpite.PlayerID)
		assert.NotEmpty(t, palpite.JogoID)
		assert.Equal(t, "j1", palpite.JogoID)
		assert.Equal(t, 1, palpite.PlacarA)
		assert.Equal(t, 0, palpite.PlacarB)
		assert.NotEmpty(t, palpite.Criado)
		assert.Empty(t, palpite.Alterado)
		assert.Equal(t, true, palpite.Status)
	})

	t.Run("Alterando status", func(t *testing.T) {
		palpite, _ := NewPalpite("p1", "j1", 1, 0)
		assert.NotNil(t, palpite)

		palpite.Disable()
		assert.Equal(t, false, palpite.Status)

		palpite.Enable()
		assert.Equal(t, true, palpite.Status)
	})

	t.Run("Criando palpite sem o jogo", func(t *testing.T) {
		palpite, err := NewPalpite("p1", "", 1, 0)
		assert.Nil(t, palpite)
		assert.NotNil(t, err)
		assert.EqualError(t, err, PalpiteEntityMsgErrorIdJogoRequerido)
	})

	t.Run("Criando palpite sem o player", func(t *testing.T) {
		palpite, err := NewPalpite("", "j1", 1, 0)
		assert.Nil(t, palpite)
		assert.NotNil(t, err)
		assert.EqualError(t, err, PalpiteEntityMsgErrorPlayerRequerido)
	})

}
