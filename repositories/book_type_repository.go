package repositories

import (
	"car-rental/models"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BookingTypeRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewBookingTypeRepository(db *gorm.DB, log *zap.Logger) *BookingTypeRepository {
	return &BookingTypeRepository{db: db, log: log}
}

func (repo *BookingTypeRepository) FindByID(id uint) (*models.BookingType, error) {
	var membership models.BookingType
	err := repo.db.First(&membership, id).Error
	if err != nil {
		repo.log.Error("Error finding membership", zap.Error(err))
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("membership not found")
		}
		return nil, err
	}
	return &membership, nil
}

func (repo *BookingTypeRepository) FindAll() ([]models.BookingType, error) {
	var memberships []models.BookingType
	err := repo.db.Find(&memberships).Error
	if err != nil {
		repo.log.Error("Error finding memberships", zap.Error(err))
		return nil, err
	}
	return memberships, nil
}
