package seeder

import (
	"car-rental/models"
)

func SeedDrivers() []models.Driver {
	return []models.Driver{
		{Name: "Stanley Baxter", NIK: "3220132938273", PhoneNumber: "81992048712", DailyCost: 150000},
		{Name: "Halsey Quinn", NIK: "3220132938293", PhoneNumber: "081992048713", DailyCost: 135000},
		{Name: "Kingsley Alvarez", NIK: "3220132938313", PhoneNumber: "081992048714", DailyCost: 150000},
		{Name: "Cecilia Flowers", NIK: "3220132938330", PhoneNumber: "081992048715", DailyCost: 155000},
		{Name: "Clarissa Brown", NIK: "3220132938351", PhoneNumber: "081992048716", DailyCost: 145000},
		{Name: "Zeph Larson", NIK: "3220132938372", PhoneNumber: "081992048717", DailyCost: 130000},
		{Name: "Zach Reynolds", NIK: "3220132938375", PhoneNumber: "081992048718", DailyCost: 140000},
	}
}
