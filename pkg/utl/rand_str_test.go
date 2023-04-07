package utl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandStr(t *testing.T) {
	t.Run("Teste gerador de string randomico", func(t *testing.T) {
		newStr := RandStr(20)
		assert.NotEmpty(t, newStr)
		assert.Len(t, newStr, 20)
	})

	t.Run("Teste do limite minimo de caracters", func(t *testing.T) {
		newStr := RandStr(-10)
		assert.NotEmpty(t, newStr)
		assert.Len(t, newStr, 1)
	})
}
