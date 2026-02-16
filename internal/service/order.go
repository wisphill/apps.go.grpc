package service

import (
	"apps.go.grpc/internal/domain"
	"apps.go.grpc/order-service/proto"
	"github.com/google/uuid"
)

type OrderService struct {
	proto.UnimplementedOrderServiceServer
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
