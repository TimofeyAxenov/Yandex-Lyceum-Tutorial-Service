package grpc

import (
	"Yandex_Lyceum_Service/pkg/api/order"
	"context"
)

type OrderService struct {
	api.UnimplementedOrderServiceServer
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *api.CreateOrderRequest) (*api.CreateOrderResponse, error) {
	return &api.CreateOrderResponse{}, nil
}
