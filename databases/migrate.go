package databases

import (
	"car-rental/models"

	"gorm.io/gorm"
)

func migrateDB(db *gorm.DB) error {
	var err error
	if err = dropTables(db); err != nil {
		return err
	}

	if err = autoMigrates(db); err != nil {
		return err
	}
	return nil
}

func autoMigrates(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Customer{},
	)
}

func dropTables(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&models.Customer{},
	)
}
