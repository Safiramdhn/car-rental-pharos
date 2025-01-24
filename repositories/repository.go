package repositories

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	Customer CustomerRepository
	Car      CarRepository
	Booking  BookingRepository
}

func NewRepository(db *gorm.DB, log *zap.Logger) Repository {
	return Repository{
		Customer: *NewCustomerRepository(db, log),
		Car:      *NewCarRepository(db, log),
		Booking:  *NewBookingRepository(db, log),
	}
}
