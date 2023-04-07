package entity

import (
	"bolao/pkg/entity"
	"errors"
	"time"
)

const (
	JogoEntityMsgErrorRodadaIdRequerida = "ID da Rodada requerida"
	JogoEntityMsgErrorTimeARequerido    = "Time A requerido"
	JogoEntityMsgErrorTimeBRequerido    = "Time B requerido"
	JogoEntityMsgErrorGolsInvalido      = "O número de gols não pode ser menor que 0"
)

type JogoEntity struct {
	ID       entity.ID `json:"id"`
	RodadaID string    `json:"rodada_id"`
	TimeAID  string    `json:"time_a"`
	GolsA    int       `json:"gols_a"`
	TimeBID  string    `json:"time_b"`
	GolsB    int       `json:"gols_b"`
	Dia      time.Time `json:"dia"`
	Local    string    `json:"local"`
	Criado   time.Time `json:"criado"`
	Alterado time.Time `json:"alterado"`
	Status   bool      `json:"status"`
}

func NewJogo(rodadaId string, timeAId string, timeBId string, dia time.Time, local string) (*JogoEntity, error) {
	jogo := &JogoEntity{
		ID:       entity.NewID(),
		RodadaID: rodadaId,
		TimeAID:  timeAId,
		GolsA:    0,
		TimeBID:  timeBId,
		GolsB:    0,
		Dia:      dia,
		Local:    local,
		Criado:   time.Now(),
		Alterado: time.Time{},
		Status:   true,
	}

	err := jogo.Validate()
	if err != nil {
		return nil, err
	}

	return jogo, nil
}

func (j *JogoEntity) Validate() error {
	if len(j.RodadaID) <= 0 {
		return errors.New(JogoEntityMsgErrorRodadaIdRequerida)
	}
	if len(j.TimeAID) <= 0 {
		return errors.New(JogoEntityMsgErrorTimeARequerido)
	}
	if len(j.TimeBID) <= 0 {
		return errors.New(JogoEntityMsgErrorTimeBRequerido)
	}
	return nil
}

func (t *JogoEntity) SetGolsTimeA(Gols int) error {
	if Gols < 0 {
		return errors.New(JogoEntityMsgErrorGolsInvalido)
	}
	t.GolsA = Gols
	return nil
}

func (t *JogoEntity) SetGolsTimeB(Gols int) error {
	if Gols < 0 {
		return errors.New(JogoEntityMsgErrorGolsInvalido)
	}
	t.GolsB = Gols
	return nil
}

func (t *JogoEntity) Enable() {
	t.Status = true
}

func (t *JogoEntity) Disable() {
	t.Status = false
}
