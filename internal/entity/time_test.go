package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTime(t *testing.T) {
	time, err := NewTime("Brasil")
	assert.Nil(t, err)
	assert.NotNil(t, time)
	assert.NotEmpty(t, time.ID)
	assert.NotEmpty(t, time.Nome)
	assert.NotEmpty(t, time.Criado)
	assert.Equal(t, "Brasil", time.Nome)
}
