package router

import (
	"fmt"
	"hallenweaver/jankenpo-go/config"
	"hallenweaver/jankenpo-go/src/game/domain"
	"hallenweaver/jankenpo-go/src/game/handler"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter configura o router Gin com suas rotas e middleware.
func SetupRouter(
	cfg *config.Config,
) *gin.Engine {
	// Definindo o modo do router Gin baseado no ambiente.
	if cfg.GoEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
		slog.Info("Aplicação Gin rodando em produção")
	} else {
		gin.SetMode(gin.DebugMode)
		slog.Info("Aplicação Gin rodando em modo debug")
	}

	// Inicializando o router Gin com middlewares
	routerEngine := gin.New()
	routerEngine.Use(gin.Logger(), gin.Recovery())

	// Endpoint de Health Check
	routerEngine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	// Inicializando handlers do jogo

	initializeGameRoutes(routerEngine)

	return routerEngine
}

func initializeGameRoutes(routerEngine *gin.Engine) {
	gameGroup := routerEngine.Group("/game")

	basicGameHandler := handler.GameHandler{
		GameType: "basic",
		ValidInputs: map[string]domain.GameInput{
			"pedra":   domain.PEDRA,
			"papel":   domain.PAPEL,
			"tesoura": domain.TESOURA,
		},
		InputAdvantages: map[domain.GameInput][]domain.GameInput{
			domain.PEDRA:   {domain.TESOURA},
			domain.PAPEL:   {domain.PEDRA},
			domain.TESOURA: {domain.PAPEL},
		},
	}

	variantGameHandler := handler.GameHandler{
		GameType: "variant",
		ValidInputs: map[string]domain.GameInput{
			"pedra":   domain.PEDRA,
			"papel":   domain.PAPEL,
			"tesoura": domain.TESOURA,
			"lagarto": domain.LAGARTO,
			"spock":   domain.SPOCK,
		},
		InputAdvantages: map[domain.GameInput][]domain.GameInput{
			domain.PEDRA:   {domain.TESOURA, domain.LAGARTO},
			domain.PAPEL:   {domain.PEDRA, domain.SPOCK},
			domain.TESOURA: {domain.PAPEL, domain.LAGARTO},
			domain.LAGARTO: {domain.PAPEL, domain.SPOCK},
			domain.SPOCK:   {domain.TESOURA, domain.PEDRA},
		},
	}

	gameHandlers := []handler.GameHandler{basicGameHandler, variantGameHandler}

	for _, gameHandler := range gameHandlers {
		fullRoute := fmt.Sprintf("/%v", gameHandler.GameType)
		gameGroup.POST(fullRoute, gameHandler.Game)
	}

}
