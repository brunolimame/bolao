package entity

import (
	"bolao/pkg/entity"
	"errors"
	"time"
)

const (
	PESO_MINIMO_RODADA              int    = 10
	PESO_INCREMENTADO_NA_RODADA     int    = 1
	MSG_ERROR_NOME_RODADA_REQUERIDO string = "Nome da rodada não definido"
	MSG_ERROR_PESO_RODADA_REQUERIDO string = "O peso não pode ser menor que 10"
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

	err := rodada.Validate()
	if err != nil {
		return nil, err
	}

	return rodada, nil
}

func (r *RodadaEntity) Validate() error {
	if len(r.Nome) <= 0 {
		return errors.New(MSG_ERROR_NOME_RODADA_REQUERIDO)
	}
	if r.Peso < PESO_MINIMO_RODADA {
		return errors.New(MSG_ERROR_PESO_RODADA_REQUERIDO)
	}
	return nil
}

func (r *RodadaEntity) Enable() {
	r.Status = true
}

func (r *RodadaEntity) Disable() {
	r.Status = false
}
