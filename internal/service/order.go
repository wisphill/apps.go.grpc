package service

import (
	"apps.go.grpc/internal/domain"
	"github.com/google/uuid"
)

type OrderService struct {
	repo repository.OrderRepository
}

func (s *OrderService) Create(userID string, amount float64) (*domain.Order, error) {
	order := &domain.Order{
		ID:     uuid.New().String(),
		UserID: userID,
		Amount: amount,
	}

	return s.repo.Save(order)
}
