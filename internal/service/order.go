package service

import (
	"apps.go.grpc/internal/domain"
	"github.com/google/uuid"
	orderV1 "github.com/wisphill/apps.api.proto/gen/service/orders/v1"
)

type OrderService struct {
	orderV1.UnimplementedOrderServiceServer
}

func NewOrderService() OrderService {
	return OrderService{}
}

func (s *OrderService) Create(userID string, amount float64) (*domain.Order, error) {
	_ = &domain.Order{
		ID:     uuid.New().String(),
		UserID: userID,
		Amount: amount,
	}

	// TODO: to be implemented
	return nil, nil
}
