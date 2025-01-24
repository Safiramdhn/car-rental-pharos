package seeder

import "car-rental/models"

func UserSeeder() []models.Customer {
	return []models.Customer{
		{Name: "Wawan Hermawan", NIK: "3372093912739", PhoneNumber: "081237123682", MembershipID: uintPtr(1)},
		{Name: "Philip Walker", NIK: "3372093912785", PhoneNumber: "081237123683", MembershipID: uintPtr(3)},
		{Name: "Hugo Fleming", NIK: "3372093912800", PhoneNumber: "081237123684", MembershipID: nil},
		{Name: "Maximillian Mendez", NIK: "3372093912848", PhoneNumber: "081237123685", MembershipID: uintPtr(2)},
		{Name: "Felix Dixon", NIK: "3372093912851", PhoneNumber: "081237123686", MembershipID: uintPtr(1)},
		{Name: "Nicholas Riddle", NIK: "3372093912929", PhoneNumber: "081237123687", MembershipID: nil},
		{Name: "Stephen Wheeler", NIK: "3372093912976", PhoneNumber: "081237123688", MembershipID: uintPtr(1)},
		{Name: "Roy Brennan", NIK: "3372093913022", PhoneNumber: "081237123689", MembershipID: nil},
		{Name: "Eliza Le", NIK: "3372093913106", PhoneNumber: "081237123690", MembershipID: nil},
		{Name: "Jesse Taylor", NIK: "3372093913126", PhoneNumber: "081237123691", MembershipID: uintPtr(3)},
		{Name: "Damien Kaufman", NIK: "3372093913202", PhoneNumber: "081237123692", MembershipID: uintPtr(1)},
		{Name: "Ayesha Richardson", NIK: "3372093913257", PhoneNumber: "081237123693", MembershipID: uintPtr(2)},
		{Name: "Margaret Stokes", NIK: "3372093913262", PhoneNumber: "081237123694", MembershipID: nil},
		{Name: "Sara Livingston", NIK: "3372093913268", PhoneNumber: "081237123695", MembershipID: nil},
		{Name: "Callie Townsend", NIK: "3372093913281", PhoneNumber: "081237123696", MembershipID: nil},
		{Name: "Lilly Fischer", NIK: "3372093913325", PhoneNumber: "081237123697", MembershipID: uintPtr(3)},
		{Name: "Theresa Barton", NIK: "3372093913335", PhoneNumber: "081237123698", MembershipID: uintPtr(1)},
		{Name: "Mia Curtis", NIK: "3372093913343", PhoneNumber: "081237123699", MembershipID: nil},
		{Name: "Flora Barlow", NIK: "3372093913400", PhoneNumber: "081237123700", MembershipID: uintPtr(2)},
		{Name: "Vanessa Patton", NIK: "3372093913434", PhoneNumber: "081237123701", MembershipID: uintPtr(2)},
	}
}

func uintPtr(value uint) *uint {
	return &value
}
