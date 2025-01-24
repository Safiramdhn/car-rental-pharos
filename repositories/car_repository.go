package repositories

import (
	"car-rental/helpers"
	"car-rental/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CarRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewCarRepository(db *gorm.DB, log *zap.Logger) *CarRepository {
	return &CarRepository{db: db, log: log}
}

func (repo *CarRepository) Create(car *models.Car) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(car).Error; err != nil {
			repo.log.Error("failed to create car", zap.Error(err))
			return err
		}
		repo.log.Info("Car created", zap.String("Name", car.Name), zap.Any("Datetime", car.CreatedAt))
		return nil
	})
}

func (repo *CarRepository) Update(car *models.Car) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(car).Error; err != nil {
			repo.log.Error("failed to update car", zap.Error(err))
			return err
		}
		repo.log.Info("Car updated", zap.String("Name", car.Name), zap.Any("Datetime", car.UpdatedAt))
		return nil
	})
}

func (repo *CarRepository) FindAll(limit, page uint) ([]models.Car, int, error) {
	var cars []models.Car
	var countData int64
	var err error

	err = repo.db.Scopes(helpers.Paginate(page, limit)).Find(&cars).Error
	if err != nil {
		repo.log.Error("Error finding cars", zap.Error(err))
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, err
	}
	err = repo.db.Model(&models.Car{}).Count(&countData).Error
	if err != nil {
		repo.log.Error("Error counting cars", zap.Error(err))
		return nil, 0, err
	}
	repo.log.Info("Cars found", zap.Any("Count", countData))
	return cars, int(countData), nil
}

func (repo *CarRepository) FindByID(id uint) (*models.Car, error) {
	var car models.Car
	err := repo.db.First(&car, id).Error
	if err != nil {
		repo.log.Error("Error finding car", zap.Error(err))
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return &car, nil
}

func (repo *CarRepository) Delete(id uint) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&models.Car{}, id).Error; err != nil {
			repo.log.Error("failed to delete car", zap.Error(err))
			return err
		}
		repo.log.Info("Car deleted", zap.Any("ID", id))
		return nil
	})
}
