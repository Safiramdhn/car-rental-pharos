package controllers

import (
	"car-rental/services"

	"go.uber.org/zap"
)

type Controller struct {
	Customer   CustomerController
	Car        CarController
	Booking    BookingController
	Membership MembershipController
}

func NewController(services services.Service, log *zap.Logger) *Controller {
	return &Controller{
		Customer:   *NewCustomerController(services.Customer, log),
		Car:        *NewCarController(services.Car, log),
		Booking:    *NewBookingController(services.Booking, log),
		Membership: *NewMembershipController(services, log),
	}
}
