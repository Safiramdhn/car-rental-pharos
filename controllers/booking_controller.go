package controllers

import (
	"car-rental/helpers"
	"car-rental/models"
	"car-rental/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BookingController struct {
	service services.BookingService
	log     *zap.Logger
}

func NewBookingController(service services.BookingService, log *zap.Logger) *BookingController {
	return &BookingController{service: service, log: log}
}

func (ctl *BookingController) CreateBooking(c *gin.Context) {
	// Get the booking request from the request body
	var booking models.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		ctl.log.Error("Error binding JSON", zap.Error(err))
		helpers.BadResponse(c, "Error binding JSON", 400)
		return
	}

	// Create the booking
	err := ctl.service.Create(&booking)
	if err != nil {
		ctl.log.Error("Error creating booking", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 500)
		return
	}

	// Return success response
	ctl.log.Info("Created booking", zap.Uint("id", booking.ID))
	helpers.GoodResponseWithData(c, "Booking has been created", 201, nil)
}

func (ctl *BookingController) UpdateBooking(c *gin.Context) {
	// Get the booking ID from the URL parameters
	idParam := c.Param("id")
	id, err := helpers.StringToUint(idParam)
	if err != nil {
		ctl.log.Error("Error converting string to uint", zap.Error(err))
		helpers.BadResponse(c, "Error converting string to uint", 400)
		return
	}

	// Get the booking request from the request body
	var booking models.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		ctl.log.Error("Error binding JSON", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 400)
		return
	}

	// Set the booking ID
	booking.ID = id
	err = ctl.service.Update(&booking)
	if err != nil {
		ctl.log.Error("Error updating booking", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 500)
		return
	}

	// Return success response
	ctl.log.Info("Updated booking", zap.Uint("id", booking.ID))
	helpers.GoodResponseWithData(c, "Booking has been updated", 200, booking)
}

func (ctl *BookingController) DeleteBooking(c *gin.Context) {
	// Get the booking ID from the URL parameters
	id := c.Param("id")
	bookingID, err := helpers.StringToUint(id)
	if err != nil {
		ctl.log.Error("Error converting string to uint", zap.Error(err))
		helpers.BadResponse(c, "Error converting string to uint", 400)
		return
	}

	// Delete the booking
	err = ctl.service.Delete(bookingID)
	if err != nil {
		ctl.log.Error("Error deleting booking", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 500)
		return
	}

	// Return success response
	ctl.log.Info("Deleted booking", zap.Uint("id", bookingID))
	helpers.GoodResponseWithData(c, "Booking has been deleted", 200, nil)
}

func (ctl *BookingController) GetBooking(c *gin.Context) {
	// Get the booking ID from the URL parameters
	id := c.Param("id")
	bookingID, err := helpers.StringToUint(id)
	if err != nil {
		ctl.log.Error("Error converting string to uint", zap.Error(err))
		helpers.BadResponse(c, "Error converting string to uint", 400)
		return
	}

	// Get the booking
	booking, err := ctl.service.Get(bookingID)
	if err != nil {
		ctl.log.Error("Error getting booking", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 500)
		return
	}

	ctl.log.Info("Successfully retrieved booking", zap.Uint("id", booking.ID))
	helpers.GoodResponseWithData(c, "Successfully retrieved booking", 200, booking)
}

func (ctl *BookingController) GetAllBookings(c *gin.Context) {
	// get limit and page parameters
	limit, _ := helpers.StringToUint(c.DefaultQuery("per_page", "10"))
	page, _ := helpers.StringToUint(c.DefaultQuery("page", "1"))

	// Get all bookings
	bookings, countData, err := ctl.service.GetAll(limit, page)
	if err != nil {
		ctl.log.Error("Error getting bookings", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 500)
		return
	}

	if bookings == nil {
		ctl.log.Info("No bookings found")
		helpers.GoodResponseWithPage(c, "No bookings found", 200, 0, 0, 0, 0, nil)
	}

	// Return success response
	ctl.log.Info("Get all bookings", zap.Int("count", countData))
	totalPage := (countData + int(limit) - 1) / int(limit) // calculate total pages correctly
	helpers.GoodResponseWithPage(c, "Cars found", 200, countData, totalPage, int(page), int(limit), bookings)
}
