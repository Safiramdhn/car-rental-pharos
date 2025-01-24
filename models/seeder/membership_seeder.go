package seeder

import "car-rental/models"

func MembershipSeeder() []models.Membership {
	return []models.Membership{
		{Name: "Bronze", Discount: 4},
		{Name: "Silver", Discount: 7},
		{Name: "Gold", Discount: 15},
	}
}
