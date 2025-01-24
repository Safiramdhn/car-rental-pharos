package seeder

import (
	"car-rental/models"
)

func BookingTypeSeeder() []models.BookingType {
	return []models.BookingType{
		{Type: "Car Only", Description: "Rent Car only"},
		{Type: "Car & Driver", Description: "Rent Car and a Driver"},
	}
}
