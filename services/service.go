package services

import (
	"car-rental/repositories"

	"go.uber.org/zap"
)

type Service struct {
	Customer CustomerService
}

func NewService(repo repositories.Repository, log *zap.Logger) *Service {
	return &Service{
		Customer: NewCustomerService(repo.Customer, log),
	}
}
