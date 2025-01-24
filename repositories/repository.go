package repositories

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	Customer        CustomerRepository
	Car             CarRepository
	Booking         BookingRepository
	Membership      MembershipRepository
	Driver          DriverRepository
	BookingType     BookingTypeRepository
	DriverIncentive DriverIncentiveRepository
}

func NewRepository(db *gorm.DB, log *zap.Logger) Repository {
	return Repository{
		Customer:        *NewCustomerRepository(db, log),
		Car:             *NewCarRepository(db, log),
		Booking:         *NewBookingRepository(db, log),
		Membership:      *NewMembershipRepository(db, log),
		Driver:          *NewDriverRepository(db, log),
		BookingType:     *NewBookingTypeRepository(db, log),
		DriverIncentive: *NewDriverIncentiveRepository(db, log),
	}
}
