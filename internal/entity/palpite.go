package entity

import (
	"bolao/pkg/entity"
	"errors"
	"time"
)

const (
	PalpiteEntityMsgErrorIdJogoRequerido = "ID do jogo é requerido"
	PalpiteEntityMsgErrorPlayerRequerido = "ID do Player é requerido"
)

type PalpiteEntity struct {
	ID       entity.ID `json:"id"`
	PlayerID string    `json:"player_id"`
	JogoID   string    `json:"jogo_id"`
	GolsA    int       `json:"gols_a"`
	GolsB    int       `json:"gols_b"`
	Pontos   int       `json:"pontos"`
	Criado   time.Time `json:"criado"`
	Alterado time.Time `json:"alterado"`
	Status   bool      `json:"status"`
}

func NewPalpite(playerId string, jogoId string) (*PalpiteEntity, error) {
	palpite := &PalpiteEntity{
		ID:       entity.NewID(),
		PlayerID: playerId,
		JogoID:   jogoId,
		GolsA:    0,
		GolsB:    0,
		Pontos:   0,
		Criado:   time.Now(),
		Alterado: time.Time{},
		Status:   true,
	}

	err := palpite.Validate()
	if err != nil {
		return nil, err
	}

	return palpite, nil
}

func (p *PalpiteEntity) Validate() error {
	if len(p.JogoID) <= 0 {
		return errors.New(PalpiteEntityMsgErrorIdJogoRequerido)
	}
	if len(p.PlayerID) <= 0 {
		return errors.New(PalpiteEntityMsgErrorPlayerRequerido)
	}
	return nil
}

func (p *PalpiteEntity) SetGols(golsTimeA int, golsTimeB int) {
	p.GolsA = golsTimeA
	p.GolsB = golsTimeB
}

func (p *PalpiteEntity) SetPontos(pontos int) {
	p.Pontos = pontos
}

func (p *PalpiteEntity) Enable() {
	p.Status = true
}

func (p *PalpiteEntity) Disable() {
	p.Status = false
}

func (p *PalpiteEntity) PontuarPalpite(GolsTimeA int, GolsTimeB int) error {

	return nil
}
