package main

import (
	"hallenweaver/jankenpo-go/config"
	"hallenweaver/jankenpo-go/pkg/infrastructure/router"
	"log/slog"
	"os"
)

func main() {
	// Olá, mundo.
	slog.Info("Inicializando aplicação jo-ken-po Itaú")

	// Inicializando variáveis de ambiente e configurações
	err := config.SetupEnv()
	if err != nil {
		slog.Info("Erro ao obter dados de ambiente", "Ambiente Usado:", os.Getenv("ENVIRONMENT"))
	}

	cfg := config.LoadConfig()

	// Inicializando o router
	appRouter := router.SetupRouter(cfg)

	// Inicializando o servidor
	config.StartServer(appRouter, cfg.AppPort)
}
