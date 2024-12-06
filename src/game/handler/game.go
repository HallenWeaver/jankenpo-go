package handler

import (
	"errors"
	"fmt"
	"hallenweaver/jankenpo-go/src/game/domain"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	GameType        string
	ValidInputs     map[string]domain.GameInput
	InputAdvantages map[domain.GameInput][]domain.GameInput
}

func (gh *GameHandler) Game(c *gin.Context) {
	var newGameRequest domain.GameRequest

	if err := c.BindJSON(&newGameRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sanitizedGameRequest, err := gh.sanitizeInput(newGameRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if sanitizedGameRequest.Jogador_1 == sanitizedGameRequest.Jogador_2 {
		c.IndentedJSON(http.StatusOK, domain.GameResponse{Resultado: "Empate!"})
		return
	}

	if gh.checkIfPlayer1Wins(sanitizedGameRequest.Jogador_1, sanitizedGameRequest.Jogador_2) {
		c.IndentedJSON(http.StatusOK, domain.GameResponse{Resultado: "Vitória do Jogador 1!"})
		return
	}

	c.IndentedJSON(http.StatusOK, domain.GameResponse{Resultado: "Vitória do Jogador 2!"})
}

func (gh *GameHandler) sanitizeInput(gameRequest domain.GameRequest) (*domain.GameRequest, error) {
	inputJogador1 := strings.ToLower(gameRequest.Jogador_1)
	inputJogador2 := strings.ToLower(gameRequest.Jogador_2)
	errMsg := ""

	if !gh.checkIfValidGameInput(inputJogador1) {
		errMsg += fmt.Sprintf("Jogador 1 usou uma opção inválida: %v |", inputJogador1)
	}
	if !gh.checkIfValidGameInput(inputJogador2) {
		errMsg += fmt.Sprintf("Jogador 2 usou uma opção inválida: %v |", inputJogador2)
	}
	if errMsg != "" {
		return nil, errors.New(errMsg)
	}

	return &domain.GameRequest{
		Jogador_1: inputJogador1,
		Jogador_2: inputJogador2,
	}, nil
}

func (gh *GameHandler) checkIfValidGameInput(playerInput string) bool {
	_, ok := gh.ValidInputs[playerInput]
	return ok
}

func (gh *GameHandler) checkIfPlayer1Wins(player1Input string, player2Input string) bool {
	winConditions := gh.InputAdvantages[domain.GameInput(player1Input)]
	for _, condition := range winConditions {
		if condition == domain.GameInput(player2Input) {
			return true
		}
	}
	return false
}
