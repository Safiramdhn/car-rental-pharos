package controllers

import (
	"car-rental/services"

	"go.uber.org/zap"
)

type Controller struct {
	Customer CustomerController
}

func NewController(services services.Service, log *zap.Logger) *Controller {
	return &Controller{
		Customer: *NewCustomerController(services.Customer, log),
	}
}
