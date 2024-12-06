package config

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// StartServer starts the HTTP server and manages graceful shutdown.
func StartServer(router *gin.Engine, port string) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		logger.Info("Iniciando servidor; ", "porta", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Falha na inicialização; ", "erro", err.Error())
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	stop()
	slog.Info("Finalizando operação do servidor")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Error("Servidor forçado a desligar;", "erro", err)
	}

	slog.Info("Finalizando execução do servidor.")
}
