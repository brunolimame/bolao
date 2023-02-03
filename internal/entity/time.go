package entity

import (
	"bolao/pkg/entity"
	"time"
)

type Time struct {
	ID     entity.ID `json:"id"`
	Nome   string    `json:"nome"`
	Criado time.Time `json:"criado"`
}

func NewTime(nome string) (*Time, error) {
	time := &Time{
		ID:     entity.NewID(),
		Nome:   nome,
		Criado: time.Now(),
	}

	return time, nil
}
