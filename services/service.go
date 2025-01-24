package services

import (
	"car-rental/repositories"

	"go.uber.org/zap"
)

type Service struct {
	Customer        CustomerService
	Car             CarService
	Booking         BookingService
	Membership      MembershipService
	Driver          DriverService
	DriverIncentive DriverIncentiveService
	BookingType     BookingTypeService
}

func NewService(repo repositories.Repository, log *zap.Logger) *Service {
	return &Service{
		Customer:        NewCustomerService(repo.Customer, log),
		Car:             NewCarService(repo.Car, log),
		Booking:         NewBookingService(repo, log),
		Membership:      NewMembershipService(repo, log),
		Driver:          NewDriverService(repo.Driver, log),
		DriverIncentive: NewDriverIncentiveService(repo.DriverIncentive, log),
		BookingType:     NewBookingTypeService(repo.BookingType, log),
	}
}
