package services

import (
	"car-rental/helpers"
	"car-rental/models"
	"car-rental/repositories"
	"errors"
	"time"

	"go.uber.org/zap"
)

type BookingService interface {
	Create(booking *models.Booking) error
	Get(bookingID uint) (*models.Booking, error)
	GetAll(limit, page uint) ([]models.Booking, int, error)
	Update(booking *models.Booking) error
	Delete(id uint) error
}

type bookingService struct {
	repo repositories.Repository
	log  *zap.Logger
}

// Create implements BookingService.
func (b *bookingService) Create(booking *models.Booking) error {
	if booking.StartRent.IsZero() || booking.EndRent.IsZero() {
		b.log.Error("Start and end date are required")
		return errors.New("start and end date are required")
	}

	if booking.CarID == 0 {
		b.log.Error("Car ID is required")
		return errors.New("car ID is required")
	}
	if booking.CustomerID == 0 {
		b.log.Error("Customer ID is required")
		return errors.New("customer ID is required")
	}

	return b.repo.Booking.Create(booking)
}

// Delete implements BookingService.
func (b *bookingService) Delete(id uint) error {
	return b.repo.Booking.Delete(id)
}

// Get implements BookingService.
func (b *bookingService) Get(bookingID uint) (*models.Booking, error) {
	return b.repo.Booking.FindByID(bookingID)
}

// GetAll implements BookingService.
func (b *bookingService) GetAll(limit uint, page uint) ([]models.Booking, int, error) {
	return b.repo.Booking.FindAll(limit, page)
}

// Update implements BookingService.
func (b *bookingService) Update(booking *models.Booking) error {
	if booking.ID == 0 {
		b.log.Error("Booking ID is required")
		return errors.New("ID is required")
	}

	oldBooking, err := b.repo.Booking.FindByID(booking.ID)
	if err != nil {
		b.log.Error("Error finding old booking", zap.Error(err))
		return err
	}

	if oldBooking.IsFinished {
		b.log.Error("Booking has already been finished", zap.Bool("isFinished", oldBooking.IsFinished))
		return errors.New("booking has already been finished")
	}

	if booking.StartRent.IsZero() || booking.EndRent.IsZero() {
		b.log.Error("Start and end date are required")
		return errors.New("start and end date are required")
	}

	if booking.CarID == 0 {
		b.log.Error("Car ID is required")
		return errors.New("car ID is required")
	}
	if booking.CustomerID == 0 {
		b.log.Error("Customer ID is required")
		return errors.New("customer ID is required")
	}

	// Set EndRent to the current date if IsFinished is updated to true
	if booking.IsFinished {
		currentDate := time.Now().Format("01/02/2006")
		booking.EndRent = helpers.FormatDate(currentDate)
	}

	return b.repo.Booking.Update(booking)
}

func NewBookingService(repo repositories.Repository, log *zap.Logger) BookingService {
	return &bookingService{repo: repo, log: log}
}
