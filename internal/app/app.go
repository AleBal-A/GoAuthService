package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
	"sso/internal/services/auth"
	"sso/internal/storage/sqlite"
	"time"
)

type Auth struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *Auth {
	// TODO: init storage
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	// TODO: init auth service (auth)
	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcAPP := grpcapp.New(log, authService, grpcPort)

	return &Auth{
		GRPCSrv: grpcAPP,
	}
}
