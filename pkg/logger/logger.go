package logger

import (
	"context"

	"go.uber.org/zap"
)

const (
	LoggerKey   = "logger"
	RequestID   = "requestID"
	ServiceName = "service"
)

type Logger interface {
	Info(ctx context.Context, msg string, fields ...zap.Field)
	Error(ctx context.Context, msg string, fields ...zap.Field)
}

type logger struct {
	serviceName string
	logger      *zap.Logger
}

func New(serviceName string) Logger {
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()
	return &logger{
		serviceName: serviceName,
		logger:      zapLogger,
	}
}

func (l logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, zap.String(ServiceName, l.serviceName), zap.String(RequestID, ctx.Value(RequestID).(string)))
	l.logger.Info(msg, fields...)
}

func (l logger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, zap.String(ServiceName, l.serviceName), zap.String(RequestID, ctx.Value(RequestID).(string)))
	l.logger.Error(msg, fields...)
}

func GetLoggerFromCtx(ctx context.Context) Logger {
	return ctx.Value(LoggerKey).(Logger)
}
