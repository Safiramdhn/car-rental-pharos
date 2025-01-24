package services

import (
	"car-rental/models"
	"car-rental/repositories"
	"errors"

	"go.uber.org/zap"
)

type DriverService interface {
	Create(driver *models.Driver) error
	Get(driverID uint) (*models.Driver, error)
	GetAll(limit, page uint) ([]models.Driver, int, error)
	Update(driver *models.Driver) error
	Delete(id uint) error
}

type driverService struct {
	repo repositories.DriverRepository
	log  *zap.Logger
}

// Create implements DriverService.
func (c *driverService) Create(driver *models.Driver) error {
	if driver.NIK == "" {
		c.log.Error("Driver NIK is required")
		return errors.New("NIK is required")
	}
	if driver.PhoneNumber == "" {
		c.log.Error("Driver Phone Number is required")
		return errors.New("phone Number is required")
	}

	if driver.DailyCost <= 0 {
		c.log.Error("Driver Daily Cost is required")
		return errors.New("daily cost is required")
	}

	c.log.Info("Creating new driver", zap.String("Name", driver.Name))
	return c.repo.Create(driver)
}

// Delete implements DriverService.
func (c *driverService) Delete(id uint) error {
	return c.repo.Delete(id)
}

// Get implements DriverService.
func (c *driverService) Get(driverID uint) (*models.Driver, error) {
	return c.repo.FindByID(driverID)
}

// GetAll implements DriverService.
func (c *driverService) GetAll(limit, page uint) ([]models.Driver, int, error) {
	return c.repo.FindAll(limit, page)
}

// Update implements DriverService.
func (c *driverService) Update(driver *models.Driver) error {
	if driver.ID <= 0 {
		c.log.Error("Driver ID is required")
		return errors.New("ID is required")
	}
	if driver.NIK == "" {
		c.log.Error("Driver NIK is required")
		return errors.New("NIK is required")
	}
	if driver.PhoneNumber == "" {
		c.log.Error("Driver Phone Number is required")
		return errors.New("phone Number is required")
	}

	if driver.DailyCost <= 0 {
		c.log.Error("Driver Daily Cost is required")
		return errors.New("daily cost is required")
	}

	c.log.Info("Updating driver", zap.Uint("ID", driver.ID))
	return c.repo.Update(driver)
}

func NewDriverService(repo repositories.DriverRepository, log *zap.Logger) DriverService {
	return &driverService{repo: repo, log: log}
}
