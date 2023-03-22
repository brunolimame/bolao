package entity

import (
	"bolao/pkg/entity"
	"time"
)

type RodadaEntity struct {
	ID       entity.ID `json:"id"`
	Nome     string    `json:"nome"`
	Peso     int       `json:"peso"`
	Criado   time.Time `json:"criado"`
	Alterado time.Time `json:"alterado"`
	Status   bool      `json:"status"`
}

func NewRodada(nome string, peso int) (*RodadaEntity, error) {
	rodada := &RodadaEntity{
		ID:       entity.NewID(),
		Nome:     nome,
		Peso:     peso,
		Criado:   time.Now(),
		Alterado: time.Time{},
		Status:   true,
	}

	return rodada, nil
}

func (r *RodadaEntity) Enable() {
	r.Status = true
}

func (r *RodadaEntity) Disable() {
	r.Status = false
}
