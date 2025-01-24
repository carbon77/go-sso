package main

import (
	"log/slog"
	"os"

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

	application.GRPCServer.MustRun()
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
