package services

import (
	"car-rental/models"
	"car-rental/repositories"
	"errors"

	"go.uber.org/zap"
)

type CarService interface {
	Create(car *models.Car) error
	Get(carID uint) (*models.Car, error)
	GetAll(limit, page uint) ([]models.Car, int, error)
	Update(car *models.Car) error
	Delete(id uint) error
}

type carService struct {
	repo repositories.CarRepository
	log  *zap.Logger
}

// Create implements CarService.
func (c *carService) Create(car *models.Car) error {
	if car.Name == "" {
		c.log.Error("Car name is required")
		return errors.New("name is required")
	}
	if car.DailyRent == 0 {
		c.log.Error("Car daily rent is required")
		return errors.New("daily rent is required")
	}
	c.log.Info("Creating new car", zap.String("Name", car.Name))
	return c.repo.Create(car)
}

// Delete implements CarService.
func (c *carService) Delete(id uint) error {
	return c.repo.Delete(id)
}

// Get implements CarService.
func (c *carService) Get(carID uint) (*models.Car, error) {
	return c.repo.FindByID(carID)
}

// GetAll implements CarService.
func (c *carService) GetAll(limit uint, page uint) ([]models.Car, int, error) {
	return c.repo.FindAll(limit, page)
}

// Update implements CarService.
func (c *carService) Update(car *models.Car) error {
	if car.ID == 0 {
		c.log.Error("Car ID is required")
		return errors.New("ID is required")
	}
	if car.Name == "" {
		c.log.Error("Car name is required")
		return errors.New("name is required")
	}
	if car.DailyRent == 0 {
		c.log.Error("Car daily rent is required")
		return errors.New("daily rent is required")
	}

	c.log.Info("Updating car", zap.String("Name", car.Name))
	return c.repo.Update(car)
}

func NewCarService(repo repositories.CarRepository, log *zap.Logger) CarService {
	return &carService{repo: repo, log: log}
}
