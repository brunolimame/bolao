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
	PlacarA  int       `json:"placar_a"`
	PlacarB  int       `json:"placar_b"`
	Criado   time.Time `json:"criado"`
	Alterado time.Time `json:"alterado"`
	Status   bool      `json:"status"`
}

func NewPalpite(playerId string, jogoId string, placarA int, placarB int) (*PalpiteEntity, error) {
	palpite := &PalpiteEntity{
		ID:       entity.NewID(),
		PlayerID: playerId,
		JogoID:   jogoId,
		PlacarA:  placarA,
		PlacarB:  placarB,
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

func (p *PalpiteEntity) Enable() {
	p.Status = true
}

func (p *PalpiteEntity) Disable() {
	p.Status = false
}
