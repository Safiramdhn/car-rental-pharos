package databases

import (
	"car-rental/config"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(cfg config.Config) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
		},
	)

	db, err := gorm.Open(postgres.Open(connectionString(cfg)), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil, err
	}

	if cfg.DB.Migrate {
		err = migrateDB(db)
		if err != nil {
			return nil, err
		}
	}
	if cfg.DB.Seeding {
		err = seedDB(db)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

func connectionString(cfg config.Config) string {
	return "host=" + cfg.DB.Host + " port=" + cfg.DB.Port + " user=" + cfg.DB.User + " dbname=" + cfg.DB.Name + " password=" + cfg.DB.Password + " sslmode=disable TimeZone=Asia/Jakarta"
}
