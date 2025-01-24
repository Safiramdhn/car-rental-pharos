package repositories

import (
	"car-rental/helpers"
	"car-rental/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BookingRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewBookingRepository(db *gorm.DB, log *zap.Logger) *BookingRepository {
	return &BookingRepository{db: db, log: log}
}

func (repo *BookingRepository) Create(booking *models.Booking) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(booking).Error; err != nil {
			repo.log.Error("failed to create booking", zap.Error(err))
			return err
		}
		repo.log.Info("Booking created", zap.Uint("Customer ID", booking.CustomerID), zap.Uint("Car ID", booking.CarID), zap.Any("Datetime", booking.CreatedAt))
		return nil
	})
}

func (repo *BookingRepository) Update(booking *models.Booking) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(booking).Error; err != nil {
			repo.log.Error("failed to update booking", zap.Error(err))
			return err
		}
		repo.log.Info("Booking updated", zap.Uint("Customer ID", booking.CustomerID), zap.Uint("Car ID", booking.CarID), zap.Any("Datetime", booking.UpdatedAt))
		return nil
	})
}

func (repo *BookingRepository) FindAll(limit, page uint) ([]models.Booking, int, error) {
	var bookings []models.Booking
	var countData int64
	var err error

	err = repo.db.Scopes(helpers.Paginate(page, limit)).Find(&bookings).Error
	if err != nil {
		repo.log.Error("Error finding bookings", zap.Error(err))
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, err
	}

	err = repo.db.Model(&models.Booking{}).Count(&countData).Error
	if err != nil {
		repo.log.Error("Error counting bookings", zap.Error(err))
		return nil, 0, err
	}

	repo.log.Info("Bookings found", zap.Any("Count", countData))
	return bookings, int(countData), nil
}

func (repo *BookingRepository) FindByID(id uint) (*models.Booking, error) {
	var booking models.Booking
	err := repo.db.First(&booking, id).Error
	if err != nil {
		repo.log.Error("Error finding booking", zap.Error(err))
	}
	return &booking, err
}

func (repo *BookingRepository) Delete(id uint) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&models.Booking{}, id).Error; err != nil {
			repo.log.Error("failed to delete booking", zap.Error(err))
			return err
		}
		repo.log.Info("Booking deleted", zap.Any("ID", id))
		return nil
	})
}
