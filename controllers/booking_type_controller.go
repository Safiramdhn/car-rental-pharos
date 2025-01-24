package controllers

import (
	"car-rental/helpers"
	"car-rental/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BookingTypeController struct {
	service services.BookingTypeService
	log     *zap.Logger
}

func NewBookingTypeController(service services.BookingTypeService, log *zap.Logger) *BookingTypeController {
	return &BookingTypeController{service: service, log: log}
}

func (ctl *BookingTypeController) GetAllBookingTypes(c *gin.Context) {
	bookingTypes, err := ctl.service.GetAll()
	if err != nil {
		ctl.log.Error("Error getting bookingTypes", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 500)
		return
	}

	if bookingTypes == nil {
		ctl.log.Info("No bookingTypes found")
		helpers.GoodResponseWithPage(c, "No bookingTypes found", 200, 0, 0, 0, 0, nil)
		return
	}

	ctl.log.Info("Get all bookingTypes")
	helpers.GoodResponseWithData(c, "BookingTypes found", 200, bookingTypes)
}
