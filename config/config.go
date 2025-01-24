package config

import (
	"flag"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppDebug bool
	Port     string
	DB       DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string

	Migrate bool
	Seeding bool
}

func NewConfig() (Config, error) {
	// set default values
	setDefaultValue()

	// set configuration path, type, and name
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("../..")
	viper.SetConfigType("dotenv")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()

	// read configuration file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return Config{}, err
	}

	// read flags
	readFlags()

	config := Config{
		AppDebug: viper.GetBool("APP_DEBUG"),
		Port:     viper.GetString("PORT"),
		DB: DBConfig{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			Name:     viper.GetString("DB_NAME"),
			Migrate:  viper.GetBool("DB_MIGRATE"),
			Seeding:  viper.GetBool("DB_SEEDING"),
		},
	}

	return config, nil
}

func setDefaultValue() {
	viper.SetDefault("APP_DEBUG", true)
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_USER", "root")
	viper.SetDefault("DB_PASSWORD", "root")
	viper.SetDefault("DB_NAME", "golang")
	viper.SetDefault("DB_MIGRATE", false)
	viper.SetDefault("DB_SEEDING", false)
}

func readFlags() {
	migrateDB := flag.Bool("m", false, "Migrate database")
	seedDB := flag.Bool("s", false, "Seed database")
	flag.Parse()
	if *migrateDB {
		viper.Set("DB_MIGRATE", true)
	}
	if *seedDB {
		viper.Set("DB_SEEDING", true)
	}
}
