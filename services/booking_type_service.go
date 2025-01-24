package services

import (
	"car-rental/models"
	"car-rental/repositories"

	"go.uber.org/zap"
)

type BookingTypeService interface {
	GetAll() ([]models.BookingType, error)
}

type bookingTypeService struct {
	repo repositories.BookingTypeRepository
	log  *zap.Logger
}

func (s *bookingTypeService) GetAll() ([]models.BookingType, error) {
	return s.repo.FindAll()
}

func NewBookingTypeService(repo repositories.BookingTypeRepository, log *zap.Logger) BookingTypeService {
	return &bookingTypeService{
		repo: repo,
		log:  log,
	}
}
