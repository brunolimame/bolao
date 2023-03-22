package entity

import (
	"bolao/pkg/entity"
	"time"
)

type Rodada struct {
	ID       entity.ID `json:"id"`
	Nome     string    `json:"nome"`
	Peso     int       `json:"peso"`
	Criado   time.Time `json:"criado"`
	Alterado time.Time `json:"alterado"`
	Status   bool      `json:"status"`
}

func NewRodada(nome string, peso int) (*Rodada, error) {
	rodada := &Rodada{
		ID:       entity.NewID(),
		Nome:     nome,
		Peso:     peso,
		Criado:   time.Now(),
		Alterado: time.Time{},
		Status:   true,
	}

	return rodada, nil
}
