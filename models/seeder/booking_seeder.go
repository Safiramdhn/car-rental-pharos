package seeder

import (
	"car-rental/helpers"
	"car-rental/models"
)

func BookingSeeder() []models.Booking {
	return []models.Booking{
		{CustomerID: 3, CarID: 2, StartRent: helpers.FormatDate("01/01/2021"), EndRent: helpers.FormatDate("01/02/2021"), TotalCost: 1000000, IsFinished: true},
		{CustomerID: 11, CarID: 2, StartRent: helpers.FormatDate("01/10/2021"), EndRent: helpers.FormatDate("01/11/2021"), TotalCost: 1000000, IsFinished: true},
		{CustomerID: 7, CarID: 1, StartRent: helpers.FormatDate("01/12/2021"), EndRent: helpers.FormatDate("01/14/2021"), TotalCost: 1500000, IsFinished: true},
		{CustomerID: 1, CarID: 15, StartRent: helpers.FormatDate("01/14/2021"), EndRent: helpers.FormatDate("01/16/2021"), TotalCost: 1800000, IsFinished: true},
		{CustomerID: 16, CarID: 7, StartRent: helpers.FormatDate("01/29/2021"), EndRent: helpers.FormatDate("01/29/2021"), TotalCost: 1000000, IsFinished: true},
		{CustomerID: 12, CarID: 14, StartRent: helpers.FormatDate("02/16/2021"), EndRent: helpers.FormatDate("02/16/2021"), TotalCost: 800000, IsFinished: true},
		{CustomerID: 5, CarID: 9, StartRent: helpers.FormatDate("02/20/2021"), EndRent: helpers.FormatDate("02/22/2021"), TotalCost: 1500000, IsFinished: true},
		{CustomerID: 2, CarID: 8, StartRent: helpers.FormatDate("03/30/2021"), EndRent: helpers.FormatDate("03/30/2021"), TotalCost: 500000, IsFinished: false},
	}
}
