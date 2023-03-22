package entity

import (
	"bolao/pkg/entity"
	"time"
)

type Time struct {
	ID       entity.ID `json:"id"`
	Nome     string    `json:"nome"`
	Escudo   string    `json:"escudo"`
	Criado   time.Time `json:"criado"`
	Alterado time.Time `json:"alterado"`
	Status   bool      `json:"status"`
}

func NewTime(nome string, escudo string) (*Time, error) {
	time := &Time{
		ID:       entity.NewID(),
		Nome:     nome,
		Escudo:   escudo,
		Criado:   time.Now(),
		Alterado: time.Time{},
		Status:   true,
	}

	return time, nil
}
