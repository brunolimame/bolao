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

	t.Run("Criando palpite sem o jogo e player", func(t *testing.T) {
		palpite, err := NewPalpite("", "")
		assert.Nil(t, palpite)
		assert.NotNil(t, err)
		assert.EqualError(t, err, PalpiteEntityMsgErrorIdJogoRequerido)
	})

	t.Run("Pontuação: Acertar o placar exato da partida", func(t *testing.T) {
		palpite, _ := NewPalpite("p1", "j1")
		assert.NotNil(t, palpite)
		assert.Equal(t, 0, palpite.Pontos)
		palpite.SetGols(5, 2)

		palpite.PontuarPalpite(1, 5, 2)
		assert.Equal(t, PalpiteEntityPontosAcertarPlacarExato, palpite.Pontos)

		palpite.PontuarPalpite(2, 5, 2)
		assert.Equal(t, PalpiteEntityPontosAcertarPlacarExato*2, palpite.Pontos)
	})
	t.Run("Pontuação: Acertar o vencedor e o número de gols da equipe vencedora", func(t *testing.T) {
		palpite, _ := NewPalpite("p1", "j1")
		assert.NotNil(t, palpite)
		assert.Equal(t, 0, palpite.Pontos)
		palpite.SetGols(5, 2)

		palpite.PontuarPalpite(1, 5, 1)
		assert.Equal(t, PalpiteEntityPontosAcertarVencedorGolsVencedor, palpite.Pontos)

		palpite.PontuarPalpite(2, 5, 1)
		assert.Equal(t, PalpiteEntityPontosAcertarVencedorGolsVencedor*2, palpite.Pontos)
	})
	t.Run("Pontuação: Acertar o vencedor e o número de gols da equipe vencedora", func(t *testing.T) {
		palpite, _ := NewPalpite("p1", "j1")
		assert.NotNil(t, palpite)
		assert.Equal(t, 0, palpite.Pontos)
		palpite.SetGols(5, 3)

		palpite.PontuarPalpite(1, 6, 4)
		assert.Equal(t, PalpiteEntityPontosAcertarVencedorDiferencaGols, palpite.Pontos)

		palpite.PontuarPalpite(2, 6, 4)
		assert.Equal(t, PalpiteEntityPontosAcertarVencedorDiferencaGols*2, palpite.Pontos)
	})
	
	t.Run("Pontuação: Acertar o vencedor e a diferença de gols entre equipes", func(t *testing.T) {
		palpite, _ := NewPalpite("p1", "j1")
		assert.NotNil(t, palpite)
		assert.Equal(t, 0, palpite.Pontos)
		palpite.SetGols(5, 3)

		palpite.PontuarPalpite(1, 6, 4)
		assert.Equal(t, PalpiteEntityPontosAcertarVencedorDiferencaGols, palpite.Pontos)

		palpite.PontuarPalpite(2, 6, 4)
		assert.Equal(t, PalpiteEntityPontosAcertarVencedorDiferencaGols*2, palpite.Pontos)
	})
	t.Run("Pontuação: Acertar o vencedor e o número de gols da equipe perdedora", func(t *testing.T) {
		palpite, _ := NewPalpite("p1", "j1")
		assert.NotNil(t, palpite)
		assert.Equal(t, 0, palpite.Pontos)
		palpite.SetGols(5, 3)

		palpite.PontuarPalpite(1, 6, 3)
		assert.Equal(t, PalpiteEntityPontosAcertarVencedorGolsPerdedor, palpite.Pontos)

		palpite.PontuarPalpite(2, 6, 3)
		assert.Equal(t, PalpiteEntityPontosAcertarVencedorGolsPerdedor*2, palpite.Pontos)
	})
	t.Run("Pontuação: Acertar que a partida terminaria em empate", func(t *testing.T) {
		palpite, _ := NewPalpite("p1", "j1")
		assert.NotNil(t, palpite)
		assert.Equal(t, 0, palpite.Pontos)
		palpite.SetGols(2, 2)

		palpite.PontuarPalpite(1, 3, 3)
		assert.Equal(t, PalpiteEntityPontosAcertarEmpate, palpite.Pontos)

		palpite.PontuarPalpite(2, 3, 3)
		assert.Equal(t, PalpiteEntityPontosAcertarEmpate*2, palpite.Pontos)
	})
	t.Run("Pontuação: Acertar apenas o vencedor da partida", func(t *testing.T) {
		palpite, _ := NewPalpite("p1", "j1")
		assert.NotNil(t, palpite)
		assert.Equal(t, 0, palpite.Pontos)
		palpite.SetGols(1, 0)

		palpite.PontuarPalpite(1, 4, 2)
		assert.Equal(t, PalpiteEntityPontosAcertarApenasVencedor, palpite.Pontos)

		palpite.PontuarPalpite(2, 4, 2)
		assert.Equal(t, PalpiteEntityPontosAcertarApenasVencedor*2, palpite.Pontos)
	})
	t.Run("Pontuação: Previu que o jogo seria um empate e não foi empate", func(t *testing.T) {
		palpite, _ := NewPalpite("p1", "j1")
		assert.NotNil(t, palpite)
		assert.Equal(t, 0, palpite.Pontos)
		palpite.SetGols(1, 1)

		palpite.PontuarPalpite(1, 2, 0)
		assert.Equal(t, PalpiteEntityPontosApostarEmpate, palpite.Pontos)

		palpite.PontuarPalpite(2, 2, 0)
		assert.Equal(t, PalpiteEntityPontosApostarEmpate*2, palpite.Pontos)
	})

}
