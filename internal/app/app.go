package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/carbon77/sso/internal/app/grpc"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	// TODO: init storage

	// TODO: init auth service

	grpcApp := grpcapp.New(log, grpcPort)
	return &App{
		GRPCServer: grpcApp,
	}
}
