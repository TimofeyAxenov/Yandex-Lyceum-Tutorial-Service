package main

import (
	"Yandex_Lyceum_Service/internal/config"
	"Yandex_Lyceum_Service/internal/transport/grpc"
	"Yandex_Lyceum_Service/pkg/db/cache"
	"Yandex_Lyceum_Service/pkg/logger"
	"context"
	"fmt"
)

const (
	serviceName = "lyceum"
)

func main() {
	ctx := context.Background()
	mainLogger := logger.New(serviceName)
	ctx = context.WithValue(ctx, logger.LoggerKey, mainLogger)
	cfg := config.New()

	if cfg == nil {
		panic("failed to load config")
	}

	redis := cache.New(cfg.RedisConfig)
	fmt.Println(redis.Ping(ctx))

	grpcserver, err := grpc.New(ctx, cfg.GRPCServerPort, cfg.RestServerPort)
	if err != nil {
		mainLogger.Error(ctx, err.Error())
		return
	}

	if err := grpcserver.Start(ctx); err != nil {
		mainLogger.Error(ctx, err.Error())
	}
}

type Hello struct {
	Field int `db:"id"`
}
