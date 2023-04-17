package entity

import (
	"bolao/pkg/entity"
	"errors"
	"time"

	"github.com/google/uuid"
)

const (
	RodadaEntityMsgErrorIdRequerido            string = "ID da rodada é requerido"
	RodadaEntityMsgErrorIdInvalido             string = "ID da rodada está inválido"
	RodadaEntityPesoMinimoRodada               int    = 10
	RodadaEntitypesoIncrementadoNaRodada       int    = 1
	RodadaEntityMsgErrorNomeRodadaRequerido    string = "Nome da rodada não definido"
	RodadaEntityMsgErrorCampeonatoIdRequerido  string = "ID do campeonato não definido"
	RodadaEntityMsgErrorPesoRodadaRequerido    string = "O peso não pode ser menor que 10"
	RodadaEntityMsgErrorJogoNaoPercenteARodada string = "O jogo não pertece a esta rodada"
)

type RodadaEntity struct {
	ID           entity.ID    `json:"id"`
	CampeonatoID string       `json:"campeonato_id"`
	Nome         string       `json:"nome"`
	Peso         int          `json:"peso"`
	Jogos        []JogoEntity `json:"jogos"`
	Criado       time.Time    `json:"criado"`
	Alterado     time.Time    `json:"alterado"`
	Status       bool         `json:"status"`
}

func NewRodada(campeaontoId, nome string, peso int) (*RodadaEntity, error) {
	rodada := &RodadaEntity{
		ID:           entity.NewID(),
		CampeonatoID: campeaontoId,
		Nome:         nome,
		Peso:         peso,
		Jogos:        []JogoEntity{},
		Criado:       time.Now(),
		Alterado:     time.Time{},
		Status:       true,
	}

	err := rodada.Validate()
	if err != nil {
		return nil, err
	}

	return rodada, nil
}

func (r *RodadaEntity) Validate() error {
	if r.ID.String() == "" || r.ID.String() == uuid.Nil.String() {
		return errors.New(RodadaEntityMsgErrorIdRequerido)
	}
	if _, err := entity.ParseID(r.ID.String()); err != nil {
		return errors.New(RodadaEntityMsgErrorIdInvalido)
	}
	if len(r.CampeonatoID) <= 0 {
		return errors.New(RodadaEntityMsgErrorCampeonatoIdRequerido)
	}
	if len(r.Nome) <= 0 {
		return errors.New(RodadaEntityMsgErrorNomeRodadaRequerido)
	}
	if r.Peso < RodadaEntityPesoMinimoRodada {
		return errors.New(RodadaEntityMsgErrorPesoRodadaRequerido)
	}
	return nil
}

func (r *RodadaEntity) AddJogo(jogo *JogoEntity) error {
	if jogo.RodadaID != r.ID.String() {
		return errors.New(RodadaEntityMsgErrorJogoNaoPercenteARodada)
	}
	r.Jogos = append(r.Jogos, *jogo)
	return nil
}

func (r *RodadaEntity) Enable() {
	r.Status = true
}

func (r *RodadaEntity) Disable() {
	r.Status = false
}
