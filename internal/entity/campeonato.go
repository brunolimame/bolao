package entity

import (
	"bolao/pkg/entity"
	"errors"
	"time"

	"github.com/google/uuid"
)

const (
	CampeonatoEntityMsgErrorIdRequerido                   string = "ID do campeonato é requerido"
	CampeonatoEntityMsgErrorIdInvalido                    string = "ID do campeonato está inválido"
	CampeonatoEntityMsgErrorNomeRequerido                 string = "Nome do campeonato não definido"
	CampeonatoEntityMsgErrorRodadaNaoPercenteAoCampeonato string = "A rodada não pertence a este campeonato"
)

type CampeonatoEntity struct {
	ID       entity.ID      `json:"id"`
	Nome     string         `json:"nome"`
	Rodadas  []RodadaEntity `json:"rodadas"`
	Criado   time.Time      `json:"criado"`
	Alterado time.Time      `json:"alterado"`
	Status   bool           `json:"status"`
}

func NewCampeonato(nome string) (*CampeonatoEntity, error) {
	campeonato := &CampeonatoEntity{
		ID:       entity.NewID(),
		Nome:     nome,
		Rodadas:  []RodadaEntity{},
		Criado:   time.Now(),
		Alterado: time.Time{},
		Status:   true,
	}

	err := campeonato.Validate()
	if err != nil {
		return nil, err
	}

	return campeonato, nil
}

func (c *CampeonatoEntity) Validate() error {
	if c.ID.String() == "" || c.ID.String() == uuid.Nil.String() {
		return errors.New(CampeonatoEntityMsgErrorIdRequerido)
	}
	if _, err := entity.ParseID(c.ID.String()); err != nil {
		return errors.New(CampeonatoEntityMsgErrorIdInvalido)
	}
	if len(c.Nome) <= 0 {
		return errors.New(CampeonatoEntityMsgErrorNomeRequerido)
	}
	return nil
}

func (c *CampeonatoEntity) AddRodada(rodada *RodadaEntity) error {
	if rodada.CampeonatoID != c.ID.String() {
		return errors.New(CampeonatoEntityMsgErrorRodadaNaoPercenteAoCampeonato)
	}
	c.Rodadas = append(c.Rodadas, *rodada)
	return nil
}

func (c *CampeonatoEntity) Enable() {
	c.Status = true
}

func (c *CampeonatoEntity) Disable() {
	c.Status = false
}
