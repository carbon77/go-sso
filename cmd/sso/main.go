package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/carbon77/sso/internal/app"
	"github.com/carbon77/sso/internal/config"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

func main() {
	config := config.MustLoad()

	log := setupLogger(config.Env)
	log.Info(
		"starting application",
		slog.String("env", config.Env),
	)

	application := app.New(
		log, config.GRPC.Port, config.StoragePath, config.TokenTTL,
	)

	go application.GRPCServer.MustRun()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop
	log.Info(
		"stopping application",
		slog.String("signal", sign.String()),
	)

	application.GRPCServer.Stop()
	log.Info("application stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envDev:
		log = slog.New(
			slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug},
			),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelInfo},
			),
		)
	}

	return log
}
