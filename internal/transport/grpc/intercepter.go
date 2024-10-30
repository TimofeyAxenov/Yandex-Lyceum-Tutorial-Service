package grpc

import (
	"Yandex_Lyceum_Service/pkg/logger"
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func LoggerIntercepter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logs := logger.GetLoggerFromCtx(ctx)
	logs.Info(ctx, "request started", zap.String("mathod", info.FullMethod))

	return handler(ctx, req)
}
