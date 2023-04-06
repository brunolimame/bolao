package entity

import (
	"bolao/pkg/entity"
	"errors"
	"time"
)

const CampeonatoEntityMsgErrorNomeRequerido = "Nome do campeonato não definido"

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
	if len(c.Nome) <= 0 {
		return errors.New(CampeonatoEntityMsgErrorNomeRequerido)
	}
	return nil
}

func (c *CampeonatoEntity) AddRodada(rodada *RodadaEntity) {
	c.Rodadas = append(c.Rodadas, *rodada)
}

func (c *CampeonatoEntity) Enable() {
	c.Status = true
}

func (c *CampeonatoEntity) Disable() {
	c.Status = false
}
