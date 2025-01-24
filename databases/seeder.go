package databases

import (
	"car-rental/models/seeder"
	"fmt"
	"reflect"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func seedDB(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// Seed the database with some data
		seeds := dataSeed()
		for _, seed := range seeds {
			err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(seed).Error
			if err != nil {
				name := reflect.TypeOf(seed).String()
				return fmt.Errorf("error seeding %s: %v", name, err)
			}
		}
		return nil
	})
}

func dataSeed() []interface{} {
	return []interface{}{
		// Define your data seeding here
		seeder.UserSeeder(),
		seeder.CarSeeder(),
		seeder.BookingSeeder(),
	}
}
