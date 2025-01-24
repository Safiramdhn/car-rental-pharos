package repositories

import (
	"car-rental/helpers"
	"car-rental/models"
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewCustomerRepository(db *gorm.DB, log *zap.Logger) *CustomerRepository {
	return &CustomerRepository{db: db, log: log}
}

func (repo *CustomerRepository) Create(customer *models.Customer) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(customer).Error; err != nil {
			repo.log.Error("failed to create customer", zap.Error(err))
			return err
		}
		repo.log.Info("Customer created", zap.String("Name", customer.Name), zap.Any("Datetime", customer.CreatedAt))
		return nil
	})
}

func (repo *CustomerRepository) FindAll(limit, page uint) ([]models.Customer, int, error) {
	var customers []models.Customer
	var countData int64

	err := repo.db.Scopes(helpers.Paginate(page, limit)).Find(&customers).Error
	if err != nil {
		repo.log.Error("Error finding customers", zap.Error(err))
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, err
	}

	err = repo.db.Model(&models.Customer{}).Count(&countData).Error
	if err != nil {
		repo.log.Error("Error counting customers", zap.Error(err))
		return nil, 0, err
	}
	repo.log.Info("Customers found", zap.Any("Count", countData))
	return customers, int(countData), nil
}

func (repo *CustomerRepository) FindByID(id uint) (*models.Customer, error) {
	var customer models.Customer
	err := repo.db.First(&customer, id).Error
	if err != nil {
		repo.log.Error("Error finding customer", zap.Error(err))
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("customer not found")
		}
		return nil, err
	}
	return &customer, nil
}

func (repo *CustomerRepository) Update(customer *models.Customer) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(customer).Error; err != nil {
			repo.log.Error("failed to update customer", zap.Error(err))
			return err
		}
		repo.log.Info("Customer updated", zap.String("Name", customer.Name), zap.Any("Datetime", customer.UpdatedAt))
		return nil
	})
}

func (repo *CustomerRepository) Delete(id uint) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&models.Customer{}).Error; err != nil {
			repo.log.Error("failed to delete customer", zap.Error(err))
			return err
		}
		repo.log.Info("Customer deleted", zap.Uint("ID", id), zap.Any("Datetime", time.Now()))
		return nil
	})
}
