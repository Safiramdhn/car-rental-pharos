package repositories

import (
	"car-rental/helpers"
	"car-rental/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DriverRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewDriverRepository(db *gorm.DB, log *zap.Logger) *DriverRepository {
	return &DriverRepository{db: db, log: log}
}

func (repo *DriverRepository) Create(driver *models.Driver) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(driver).Error; err != nil {
			repo.log.Error("failed to create driver", zap.Error(err))
			return err
		}
		repo.log.Info("Customer created", zap.String("Name", driver.Name), zap.Any("Datetime", driver.CreatedAt))
		return nil
	})
}

func (repo *DriverRepository) FindAll(limit, page uint) ([]models.Driver, int, error) {
	var drivers []models.Driver
	var countData int64

	err := repo.db.Scopes(helpers.Paginate(page, limit)).Find(&drivers).Error
	if err != nil {
		repo.log.Error("failed to find all drivers", zap.Error(err))
		return nil, 0, err
	}

	err = repo.db.Model(&models.Driver{}).Count(&countData).Error
	if err != nil {
		repo.log.Error("failed to count drivers", zap.Error(err))
		return nil, 0, err
	}

	repo.log.Info("Drivers found", zap.Any("Count", countData))
	return drivers, int(countData), nil
}

func (repo *DriverRepository) FindByID(id uint) (*models.Driver, error) {
	var driver models.Driver
	err := repo.db.First(&driver, id).Error
	if err != nil {
		repo.log.Error("failed to find driver", zap.Error(err))
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return &driver, nil
}

func (repo *DriverRepository) Update(driver *models.Driver) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(driver).Error; err != nil {
			repo.log.Error("failed to update driver", zap.Error(err))
			return err
		}
		repo.log.Info("Customer updated", zap.String("Name", driver.Name), zap.Any("Datetime", driver.UpdatedAt))
		return nil
	})
}

func (repo *DriverRepository) Delete(id uint) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&models.Driver{}, id).Error; err != nil {
			repo.log.Error("failed to delete driver", zap.Error(err))
			return err
		}
		repo.log.Info("Driver deleted", zap.Uint("ID", id))
		return nil
	})
}
