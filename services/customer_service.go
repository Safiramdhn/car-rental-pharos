package services

import (
	"car-rental/models"
	"car-rental/repositories"
	"errors"

	"go.uber.org/zap"
)

type CustomerService interface {
	Create(customer *models.Customer) error
	Get(customerID uint) (*models.Customer, error)
	GetAll(limit, page uint) ([]models.Customer, int, error)
	Update(customer *models.Customer) error
	Delete(id uint) error
}

type customerService struct {
	repo repositories.CustomerRepository
	log  *zap.Logger
}

// Create implements CustomerService.
func (c *customerService) Create(customer *models.Customer) error {
	if customer.NIK == "" {
		c.log.Error("Customer NIK is required")
		return errors.New("NIK is required")
	}
	if customer.PhoneNumber == "" {
		c.log.Error("Customer Phone Number is required")
		return errors.New("phone Number is required")
	}
	c.log.Info("Creating new customer", zap.String("Name", customer.Name))
	return c.repo.Create(customer)
}

// Delete implements CustomerService.
func (c *customerService) Delete(id uint) error {
	return c.repo.Delete(id)
}

// Get implements CustomerService.
func (c *customerService) Get(customerID uint) (*models.Customer, error) {
	return c.repo.FindByID(customerID)
}

// GetAll implements CustomerService.
func (c *customerService) GetAll(limit, page uint) ([]models.Customer, int, error) {
	return c.repo.FindAll(limit, page)
}

// Update implements CustomerService.
func (c *customerService) Update(customer *models.Customer) error {
	return c.repo.Update(customer)
}

func NewCustomerService(repo repositories.CustomerRepository, log *zap.Logger) CustomerService {
	return &customerService{repo: repo, log: log}
}
