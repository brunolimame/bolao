package entity

import (
	"bolao/pkg/entity"
	"errors"
	"time"
)

const MSG_ERROR_NOME_TIME_REQUERIDO = "Nome do time não definido"

type TimeEntity struct {
	ID       entity.ID `json:"id"`
	Nome     string    `json:"nome"`
	Escudo   string    `json:"escudo"`
	Criado   time.Time `json:"criado"`
	Alterado time.Time `json:"alterado"`
	Status   bool      `json:"status"`
}

func NewTime(nome string, escudo string) (*TimeEntity, error) {
	time := &TimeEntity{
		ID:       entity.NewID(),
		Nome:     nome,
		Escudo:   escudo,
		Criado:   time.Now(),
		Alterado: time.Time{},
		Status:   true,
	}

	err := time.Validate()
	if err != nil {
		return nil, err
	}

	return time, nil
}

func (r *TimeEntity) Validate() error {
	if len(r.Nome) <= 0 {
		return errors.New(MSG_ERROR_NOME_TIME_REQUERIDO)
	}
	return nil
}

func (t *TimeEntity) Enable() {
	t.Status = true
}

func (t *TimeEntity) Disable() {
	t.Status = false
}
