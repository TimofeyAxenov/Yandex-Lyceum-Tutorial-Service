package grpc

import (
	"Yandex_Lyceum_Service/pkg/api/order"
	"Yandex_Lyceum_Service/pkg/logger"
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"log"

	"fmt"
)

type Server struct {
	grpcServer *grpc.Server
	restServer *http.Server
	listener   net.Listener
}

func New(ctx context.Context, port, restPort int) (*Server, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(LoggerIntercepter),
	}
	grpcServer := grpc.NewServer(opts...)

	restSrv := runtime.NewServeMux()
	if err = api.RegisterOrderServiceHandlerServer(ctx, restSrv, NewOrderService()); err != nil {
		return nil, err
	}

	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%d", restPort),
		Handler: restSrv,
	}

	api.RegisterOrderServiceServer(grpcServer, NewOrderService())

	return &Server{grpcServer: grpcServer, listener: lis, restServer: httpServer}, nil
}

func (s *Server) Start(ctx context.Context) error {
	eg := errgroup.Group{}

	eg.Go(func() error {
		logger.GetLoggerFromCtx(ctx).Info(ctx, "starting gRPC server", zap.Int("port", s.listener.Addr().(*net.TCPAddr).Port))
		return s.grpcServer.Serve(s.listener)
	})

	eg.Go(func() error {
		logger.GetLoggerFromCtx(ctx).Info(ctx, "starting rest server", zap.String("port", s.restServer.Addr))
		return s.restServer.ListenAndServe()
	})

	return eg.Wait()
}
