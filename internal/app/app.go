package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
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

	// TODO: init auth service (auth)

	grpcAPP := grpcapp.New(log, grpcPort)

	return &Auth{
		GRPCSrv: grpcAPP,
	}
}
