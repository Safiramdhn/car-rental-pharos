package repositories

import (
	"car-rental/helpers"
	"car-rental/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DriverIncentiveRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewDriverIncentiveRepository(db *gorm.DB, log *zap.Logger) *DriverIncentiveRepository {
	return &DriverIncentiveRepository{db: db, log: log}
}

func (repo *DriverIncentiveRepository) FindAll(limit, page uint) ([]models.DriverIncentive, int, error) {
	var customers []models.DriverIncentive
	var countData int64

	err := repo.db.Scopes(helpers.Paginate(page, limit)).Find(&customers).Error
	if err != nil {
		repo.log.Error("Error finding customers", zap.Error(err))
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, err
	}

	err = repo.db.Model(&models.DriverIncentive{}).Count(&countData).Error
	if err != nil {
		repo.log.Error("Error counting customers", zap.Error(err))
		return nil, 0, err
	}
	repo.log.Info("DriverIncentives found", zap.Any("Count", countData))
	return customers, int(countData), nil
}
