package service

import (
	"context"

	"apps.go.grpc/internal/domain"
	"github.com/google/uuid"
	orderV1 "github.com/wisphill/apps.api.proto/gen/service/orders/v1"
)

type OrderService struct {
	orderV1.UnimplementedOrderServiceServer
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *orderV1.CreateOrderRequest) (*orderV1.CreateOrderResponse, error) {
	_ = &domain.Order{
		ID:     uuid.New().String(),
		UserID: in.GetUserId(),
		Amount: in.GetAmount(),
	}

	// TODO: to be implemented
	return nil, nil
}
