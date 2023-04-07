package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJogo(t *testing.T) {

	t.Run("Criando novo jogo", func(t *testing.T) {
		jogo, err := NewJogo("r1", "j1", "j2", time.Now(), "Campo")
		assert.Nil(t, err)
		assert.NotNil(t, jogo)
		assert.NotEmpty(t, jogo.ID)
		assert.NotEmpty(t, jogo.RodadaID)
		assert.Equal(t, "r1", jogo.RodadaID)
		assert.NotEmpty(t, jogo.TimeAID)
		assert.Equal(t, "j1", jogo.TimeAID)
		assert.NotEmpty(t, jogo.TimeBID)
		assert.Equal(t, "j2", jogo.TimeBID)
		assert.NotEmpty(t, jogo.Criado)
		assert.Empty(t, jogo.Alterado)
		assert.Equal(t, true, jogo.Status)
	})

	t.Run("Alterando status", func(t *testing.T) {
		jogo, _ := NewJogo("r1", "j1", "j2", time.Now(), "Campo")
		assert.NotNil(t, jogo)

		jogo.Disable()
		assert.Equal(t, false, jogo.Status)

		jogo.Enable()
		assert.Equal(t, true, jogo.Status)
	})

	t.Run("Criando jogo sem a rodada", func(t *testing.T) {
		jogo, err := NewJogo("", "", "", time.Now(), "Campo")
		assert.Nil(t, jogo)
		assert.NotNil(t, err)
		assert.EqualError(t, err, JogoEntityMsgErrorRodadaIdRequerida)
	})

	t.Run("Criando jogo sem o time A", func(t *testing.T) {
		jogo, err := NewJogo("r1", "", "t2", time.Now(), "Campo")
		assert.Nil(t, jogo)
		assert.NotNil(t, err)
		assert.EqualError(t, err, JogoEntityMsgErrorTimeARequerido)
	})

	t.Run("Criando jogo sem o time B", func(t *testing.T) {
		jogo, err := NewJogo("r1", "t1", "", time.Now(), "Campo")
		assert.Nil(t, jogo)
		assert.NotNil(t, err)
		assert.EqualError(t, err, JogoEntityMsgErrorTimeBRequerido)
	})
	t.Run("Setando número de gols", func(t *testing.T) {
		jogo, err := NewJogo("r1", "t1", "t2", time.Now(), "Campo")
		assert.Nil(t, err)
		assert.NotNil(t, jogo)
		err = jogo.SetGolsTimeA(10)
		assert.Nil(t, err)
		assert.Equal(t, 10, jogo.GolsA)
		err = jogo.SetGolsTimeB(5)
		assert.Nil(t, err)
		assert.Equal(t, 5, jogo.GolsB)
	})
	t.Run("Setando número de gols inválidos", func(t *testing.T) {
		jogo, err := NewJogo("r1", "t1", "t2", time.Now(), "Campo")
		assert.Nil(t, err)
		assert.NotNil(t, jogo)
		err = jogo.SetGolsTimeA(-1)
		assert.NotNil(t, err)
		assert.EqualError(t, err, JogoEntityMsgErrorGolsInvalido)
		err = jogo.SetGolsTimeB(-1)
		assert.NotNil(t, err)
		assert.EqualError(t, err, JogoEntityMsgErrorGolsInvalido)
	})
}
