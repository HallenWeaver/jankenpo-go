package domain

type GameRequest struct {
	Jogador_1 string `json:"jogador_1" binding:"required"`
	Jogador_2 string `json:"jogador_2" binding:"required"`
}

type GameResponse struct {
	Resultado string `json:"resultado"`
}

type GameInput string

const (
	LAGARTO = GameInput("lagarto")
	PAPEL   = GameInput("papel")
	PEDRA   = GameInput("pedra")
	SPOCK   = GameInput("spock")
	TESOURA = GameInput("tesoura")
)
