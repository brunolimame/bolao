package entity

import (
	"testing"
	"time"

	"github.com/google/uuid"
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

	t.Run("Testando ID", func(t *testing.T) {
		palpite := &PalpiteEntity{
			ID:       uuid.UUID{},
			PlayerID: "p1",
			JogoID:   "p2",
			GolsA:    0,
			GolsB:    0,
			Pontos:   0,
			Criado:   time.Time{},
			Alterado: time.Time{},
			Status:   true,
		}

		assert.Empty(t, palpite.ID)
		err := palpite.Validate()
		assert.NotNil(t, err)
		assert.EqualError(t, err, PalpiteEntityMsgErrorIdRequerido)
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

	t.Run("Pontuação: Testar Placar negativo", func(t *testing.T) {
		palpite, _ := NewPalpite("p1", "j1")
		assert.NotNil(t, palpite)
		assert.Equal(t, 0, palpite.Pontos)
		err := palpite.SetGols(-1, 2)
		assert.NotNil(t, err)
		assert.EqualError(t, err, PalpiteEntityMsgErrorNumeroGolsPalpiteNegativo)

		err = palpite.SetGols(0, -2)
		assert.NotNil(t, err)
		assert.EqualError(t, err, PalpiteEntityMsgErrorNumeroGolsPalpiteNegativo)

		palpite.SetGols(1, 0)
		err = palpite.PontuarPalpite(1, -1, 2)
		assert.NotNil(t, err)
		assert.EqualError(t, err, PalpiteEntityMsgErrorNumeroGolsJogoNegativo)

		err = palpite.PontuarPalpite(1, 0, -1)
		assert.NotNil(t, err)
		assert.EqualError(t, err, PalpiteEntityMsgErrorNumeroGolsJogoNegativo)
	})

	t.Run("Pontuação: Acertar o placar exato da partida", func(t *testing.T) {
		palpite, _ := NewPalpite("p1", "j1")
		assert.NotNil(t, palpite)
		assert.Equal(t, 0, palpite.Pontos)
		err := palpite.SetGols(5, 2)
		assert.Nil(t, err)

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
	t.Run("Pontuação: 0 pontos", func(t *testing.T) {
		palpite, _ := NewPalpite("p1", "j1")
		assert.NotNil(t, palpite)
		assert.Equal(t, 0, palpite.Pontos)
		palpite.SetGols(2, 0)

		palpite.PontuarPalpite(1, 1, 5)
		assert.Equal(t, 0, palpite.Pontos)

		palpite.PontuarPalpite(2, 1, 5)
		assert.Equal(t, 0, palpite.Pontos)
	})

}

func BenchmarkPontuarPalpite(b *testing.B) {
	palpite, _ := NewPalpite("p1", "j1")
	_ = palpite.SetGols(2, 0)
	for i := 0; i < b.N; i++ {
		_ = palpite.PontuarPalpite(1, 0, 5)
	}
}
