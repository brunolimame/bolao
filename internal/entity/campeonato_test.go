package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCampeonato(t *testing.T) {
	t.Run("Criando novo campeonato", func(t *testing.T) {
		campeonato, err := NewCampeonato("Campeonato Brasileiro")
		assert.Nil(t, err)
		assert.NotNil(t, campeonato)
		assert.NotEmpty(t, campeonato.ID)
		assert.NotEmpty(t, campeonato.Nome)
		assert.Equal(t, "Campeonato Brasileiro", campeonato.Nome)
		assert.Empty(t, campeonato.Alterado)
		assert.NotEmpty(t, campeonato.Criado)
		assert.Empty(t, campeonato.Alterado)
		assert.Equal(t, true, campeonato.Status)
	})
	t.Run("Alterando status", func(t *testing.T) {
		campeonato, _ := NewCampeonato("Campeonato Brasileiro")
		assert.NotNil(t, campeonato)

		campeonato.Disable()
		assert.Equal(t, false, campeonato.Status)

		campeonato.Enable()
		assert.Equal(t, true, campeonato.Status)
	})

	t.Run("Adicionar rodada", func(t *testing.T) {
		campeonato, _ := NewCampeonato("Campeonato Brasileiro")
		assert.NotNil(t, campeonato)
		rodada1, _ := NewRodada("Rodada 01", 10)
		rodada2, _ := NewRodada("Rodada 02", 10)
		campeonato.AddRodada(rodada1)
		campeonato.AddRodada(rodada2)
		assert.Equal(t, 2, len(campeonato.Rodada))
	})

	t.Run("Criando campeonato com nome em branco", func(t *testing.T) {
		campeonato, err := NewCampeonato("")
		assert.Nil(t, campeonato)
		assert.NotNil(t, err)
		assert.EqualError(t, err, MSG_ERROR_NOME_CAMPEONATO_REQUERIDO)
	})
}
