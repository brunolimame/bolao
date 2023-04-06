package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRodada(t *testing.T) {

	t.Run("Criando nova Rodada", func(t *testing.T) {
		rodada, err := NewRodada("c1", "Rodada 01", 10)
		assert.Nil(t, err)
		assert.NotNil(t, rodada)
		assert.NotEmpty(t, rodada.ID)
		assert.NotEmpty(t, rodada.CampeonatoID)
		assert.Equal(t, "c1", rodada.CampeonatoID)
		assert.NotEmpty(t, rodada.Nome)
		assert.Equal(t, "Rodada 01", rodada.Nome)
		assert.Equal(t, 10, rodada.Peso)
		assert.NotEmpty(t, rodada.Criado)
		assert.Empty(t, rodada.Alterado)
		assert.Equal(t, true, rodada.Status)
	})

	t.Run("Alterando status", func(t *testing.T) {
		rodada, _ := NewRodada("c1", "Rodada 01", 10)
		assert.NotNil(t, rodada)

		rodada.Disable()
		assert.Equal(t, false, rodada.Status)

		rodada.Enable()
		assert.Equal(t, true, rodada.Status)
	})

	t.Run("Adicionar jogos", func(t *testing.T) {
		rodada, _ := NewRodada("c1", "Rodada 01", 10)
		assert.NotNil(t, rodada)
		jogo1, _ := NewJogo(rodada.ID.String(), "ta1ID", "tb1ID", time.Now(), "Campo")
		jogo2, _ := NewJogo(rodada.ID.String(), "ta2ID", "tb2ID", time.Now(), "Campo")
		err := rodada.AddJogo(jogo1)
		assert.Nil(t, err)
		rodada.AddJogo(jogo2)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(rodada.Jogos))
	})
	t.Run("Não aceitar adicionar jogos de outra rodada", func(t *testing.T) {
		rodada, _ := NewRodada("c1", "Rodada 01", 10)
		assert.NotNil(t, rodada)
		jogo1, _ := NewJogo(rodada.ID.String(), "ta1ID", "tb1ID", time.Now(), "Campo")
		jogo2, _ := NewJogo("rid", "ta2ID", "tb2ID", time.Now(), "Campo")
		err := rodada.AddJogo(jogo1)
		assert.Equal(t, 1, len(rodada.Jogos))
		err = rodada.AddJogo(jogo2)
		assert.EqualError(t, err, RodadaEntityMsgErrorJogoNaoPercenteARodada)
	})

	t.Run("Criando rodada com campeonato em branco", func(t *testing.T) {
		rodada, err := NewRodada("", "Rodada 01", 10)
		assert.Nil(t, rodada)
		assert.NotNil(t, err)
		assert.EqualError(t, err, RodadaEntityMsgErrorCampeonatoIdRequerido)
	})
	t.Run("Criando rodada com nome em branco", func(t *testing.T) {
		rodada, err := NewRodada("c1", "", 10)
		assert.Nil(t, rodada)
		assert.NotNil(t, err)
		assert.EqualError(t, err, RodadaEntityMsgErrorNomeRodadaRequerido)
	})

	t.Run("Criando rodada com peso abaixo do minimo", func(t *testing.T) {
		rodada, err := NewRodada("c1", "Rodada 1", 8)
		assert.Nil(t, rodada)
		assert.NotNil(t, err)
		assert.EqualError(t, err, RodadaEntityMsgErrorPesoRodadaRequerido)
	})

}
