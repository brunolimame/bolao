package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalpite(t *testing.T) {

	t.Run("Criando novo palpite", func(t *testing.T) {
		palpite, err := NewPalpite("p1", "j1")
		assert.Nil(t, err)
		assert.NotNil(t, palpite)
		assert.NotEmpty(t, palpite.ID)
		assert.NotEmpty(t, palpite.PlayerID)
		assert.Equal(t, "p1", palpite.PlayerID)
		assert.NotEmpty(t, palpite.JogoID)
		assert.Equal(t, "j1", palpite.JogoID)
		assert.Equal(t, 0, palpite.GolsA)
		assert.Equal(t, 0, palpite.GolsB)
		assert.Equal(t, 0, palpite.Pontos)
		assert.NotEmpty(t, palpite.Criado)
		assert.Empty(t, palpite.Alterado)
		assert.Equal(t, true, palpite.Status)
	})

	t.Run("Alterando status", func(t *testing.T) {
		palpite, _ := NewPalpite("p1", "j1")
		assert.NotNil(t, palpite)

		palpite.Disable()
		assert.Equal(t, false, palpite.Status)

		palpite.Enable()
		assert.Equal(t, true, palpite.Status)
	})

	t.Run("Adicionando Pontos", func(t *testing.T) {
		palpite, _ := NewPalpite("p1", "j1")
		assert.NotNil(t, palpite)
		assert.Equal(t, 0, palpite.Pontos)
		palpite.SetPontos(10)
		assert.Equal(t, 10, palpite.Pontos)
	})
	t.Run("Adicionando placar", func(t *testing.T) {
		palpite, _ := NewPalpite("p1", "j1")
		assert.NotNil(t, palpite)
		assert.Equal(t, 0, palpite.GolsA)
		assert.Equal(t, 0, palpite.GolsB)
		palpite.SetGols(5, 2)
		assert.Equal(t, 5, palpite.GolsA)
		assert.Equal(t, 2, palpite.GolsB)
	})

	t.Run("Criando palpite sem o jogo", func(t *testing.T) {
		palpite, err := NewPalpite("p1", "")
		assert.Nil(t, palpite)
		assert.NotNil(t, err)
		assert.EqualError(t, err, PalpiteEntityMsgErrorIdJogoRequerido)
	})

	t.Run("Criando palpite sem o player", func(t *testing.T) {
		palpite, err := NewPalpite("", "j1")
		assert.Nil(t, palpite)
		assert.NotNil(t, err)
		assert.EqualError(t, err, PalpiteEntityMsgErrorPlayerRequerido)
	})

}
