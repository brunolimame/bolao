package entity

import (
	"bolao/pkg/entity"
	"errors"
	"time"

	"github.com/google/uuid"
)

const (
	TimeEntityMsgErrorIdRequerido       string = "ID do time é requerido"
	TimeEntityMsgErrorIdInvalido        string = "ID do time está inválido"
	TimeEntityMsgErrorNomeTimeRequerido string = "Nome do time não definido"
)

type TimeEntity struct {
	ID       entity.ID `json:"id"`
	Nome     string    `json:"nome"`
	Escudo   string    `json:"escudo"`
	Criado   time.Time `json:"criado"`
	Alterado time.Time `json:"alterado"`
	Status   bool      `json:"status"`
}

func NewTime(nome, escudo string) (*TimeEntity, error) {
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

func (t *TimeEntity) Validate() error {
	if t.ID.String() == "" || t.ID.String() == uuid.Nil.String() {
		return errors.New(TimeEntityMsgErrorIdRequerido)
	}
	if _, err := entity.ParseID(t.ID.String()); err != nil {
		return errors.New(TimeEntityMsgErrorIdInvalido)
	}
	if len(t.Nome) <= 0 {
		return errors.New(TimeEntityMsgErrorNomeTimeRequerido)
	}
	return nil
}

func (t *TimeEntity) Enable() {
	t.Status = true
}

func (t *TimeEntity) Disable() {
	t.Status = false
}
