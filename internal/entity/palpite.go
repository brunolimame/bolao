package entity

import (
	"bolao/pkg/entity"
	"errors"
	"math"
	"time"

	"github.com/google/uuid"
)

const (
	PalpiteEntityMsgErrorIdRequerido                string = "ID do palpite é requerido"
	PalpiteEntityMsgErrorIdInvalido                 string = "ID do palpite está inválido"
	PalpiteEntityMsgErrorIdJogoRequerido            string = "ID do jogo é requerido"
	PalpiteEntityMsgErrorPlayerRequerido            string = "ID do Player é requerido"
	PalpiteEntityPontosAcertarPlacarExato           int    = 25
	PalpiteEntityPontosAcertarVencedorGolsVencedor  int    = 20
	PalpiteEntityPontosAcertarVencedorDiferencaGols int    = 15
	PalpiteEntityPontosAcertarVencedorGolsPerdedor  int    = 12
	PalpiteEntityPontosAcertarEmpate                int    = 8
	PalpiteEntityPontosAcertarApenasVencedor        int    = 4
	PalpiteEntityPontosApostarEmpate                int    = 2
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

func NewPalpite(playerId, jogoId string) (*PalpiteEntity, error) {
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
	if p.ID.String() == "" || p.ID.String() == uuid.Nil.String() {
		return errors.New(PalpiteEntityMsgErrorIdRequerido)
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return errors.New(PalpiteEntityMsgErrorIdInvalido)
	}
	if len(p.JogoID) <= 0 {
		return errors.New(PalpiteEntityMsgErrorIdJogoRequerido)
	}
	if len(p.PlayerID) <= 0 {
		return errors.New(PalpiteEntityMsgErrorPlayerRequerido)
	}
	return nil
}

func (p *PalpiteEntity) SetGols(golsTimeA, golsTimeB int) {
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

func (p *PalpiteEntity) PontuarPalpite(PesoRodada, JogoTimeA, JogoTimeB int) {

	pontuacao := 0
	palpiteTimeA := p.GolsA
	palpiteTimeB := p.GolsB
	diferencaGols := int(math.Abs(float64(JogoTimeA-JogoTimeB))) - int(math.Abs(float64(palpiteTimeA-palpiteTimeB)))

	if palpiteTimeA == JogoTimeA && palpiteTimeB == JogoTimeB { // Acertar o placar exato da partida, 25 pontos
		pontuacao = PalpiteEntityPontosAcertarPlacarExato * PesoRodada
	} else if (JogoTimeA > JogoTimeB && palpiteTimeA > palpiteTimeB) || (JogoTimeA < JogoTimeB && palpiteTimeA < palpiteTimeB) {
		// Acertou o vencedor
		if palpiteTimeA == JogoTimeA { // Acertar o vencedor e o número de gols da equipe vencedora
			pontuacao = PalpiteEntityPontosAcertarVencedorGolsVencedor * PesoRodada
		} else if diferencaGols == 0 { // Acertar o vencedor e a diferença de gols entre o equipes
			pontuacao = PalpiteEntityPontosAcertarVencedorDiferencaGols * PesoRodada
		} else if (JogoTimeA > JogoTimeB && palpiteTimeB == JogoTimeB) || (JogoTimeA < JogoTimeB && palpiteTimeA == JogoTimeA) {
			pontuacao = PalpiteEntityPontosAcertarVencedorGolsPerdedor * PesoRodada // Acertar o vencedor e o número de gols da equipe perdedora
		} else {
			pontuacao = PalpiteEntityPontosAcertarApenasVencedor * PesoRodada // Acertar apenas o vencedor da partida
		}
	} else if palpiteTimeA == palpiteTimeB && JogoTimeA == JogoTimeB {
		pontuacao = PalpiteEntityPontosAcertarEmpate * PesoRodada // Acertar que a partida terminaria em empate
	} else if palpiteTimeA == palpiteTimeB && JogoTimeA != JogoTimeB {

		pontuacao = PalpiteEntityPontosApostarEmpate * PesoRodada // Previu que o jogo seria um empate e não foi empate
	} else {
		pontuacao = 0
	}

	p.Pontos = pontuacao
}
