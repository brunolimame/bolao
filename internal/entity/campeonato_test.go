package entity

import (
	"testing"
	"time"

	"github.com/google/uuid"
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
		rodada1, _ := NewRodada(campeonato.ID.String(), "Rodada 01", 10)
		rodada2, _ := NewRodada(campeonato.ID.String(), "Rodada 02", 10)
		campeonato.AddRodada(rodada1)
		campeonato.AddRodada(rodada2)
		assert.Equal(t, 2, len(campeonato.Rodadas))
	})
	t.Run("Testando ID", func(t *testing.T) {
		campeonato := &CampeonatoEntity{
			ID:       uuid.UUID{},
			Nome:     "C1",
			Rodadas:  []RodadaEntity{},
			Criado:   time.Now(),
			Alterado: time.Time{},
			Status:   true,
		}

		assert.Empty(t, campeonato.ID)
		err := campeonato.Validate()
		assert.NotNil(t, err)
		assert.EqualError(t, err, CampeonatoEntityMsgErrorIdRequerido)
	})

	t.Run("Não aceitar adicionar rodadas de outro campeonato", func(t *testing.T) {
		campeonato, _ := NewCampeonato("Campeonato Brasileiro")
		assert.NotNil(t, campeonato)
		rodada1, _ := NewRodada(campeonato.ID.String(), "Rodada 01", 10)
		rodada2, _ := NewRodada("c1", "Rodada 02", 10)
		err := campeonato.AddRodada(rodada1)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(campeonato.Rodadas))
		err = campeonato.AddRodada(rodada2)
		assert.NotNil(t, err)
		assert.EqualError(t, err, CampeonatoEntityMsgErrorRodadaNaoPercenteAoCampeonato)

	})

	t.Run("Criando campeonato com nome em branco", func(t *testing.T) {
		campeonato, err := NewCampeonato("")
		assert.Nil(t, campeonato)
		assert.NotNil(t, err)
		assert.EqualError(t, err, CampeonatoEntityMsgErrorNomeRequerido)
	})
}
