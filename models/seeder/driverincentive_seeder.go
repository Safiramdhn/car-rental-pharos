package seeder

import (
	"car-rental/models"
)

func SeedDriverIncentives() []models.DriverIncentive {
	return []models.DriverIncentive{
		{BookingID: 6, Incentive: 40000},
		{BookingID: 7, Incentive: 75000},
		{BookingID: 8, Incentive: 25000},
	}
}
