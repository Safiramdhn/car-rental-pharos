package services

import (
	"car-rental/models"
	"car-rental/repositories"

	"go.uber.org/zap"
)

type DriverIncentiveService interface {
	GetAll(limit, page uint) ([]models.DriverIncentive, int, error)
}

type driverIncentiveService struct {
	repo repositories.DriverIncentiveRepository
	log  *zap.Logger
}

func (s *driverIncentiveService) GetAll(limit, page uint) ([]models.DriverIncentive, int, error) {
	return s.repo.FindAll(limit, page)
}

func NewDriverIncentiveService(repo repositories.DriverIncentiveRepository, log *zap.Logger) DriverIncentiveService {
	return &driverIncentiveService{
		repo: repo,
		log:  log,
	}
}
