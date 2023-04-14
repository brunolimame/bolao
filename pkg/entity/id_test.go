package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestEntityId(t *testing.T) {
	t.Run("Teste gerador de UUID", func(t *testing.T) {
		newUuid := NewID()
		assert.NotEmpty(t, newUuid.String())
	})

	t.Run("Teste validação do UUID", func(t *testing.T) {
		_, err := ParseID(uuid.New().String())
		assert.Nil(t, err)

		_, err = ParseID("123")
		assert.NotNil(t, err)
	})
}
