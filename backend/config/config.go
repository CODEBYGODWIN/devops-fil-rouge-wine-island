package config

import (
	"api-go/database/dbmodel"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	ProductEntryRepository dbmodel.ProductRepository
}

func New() (*Config, error) {
	config := Config{}
	var err error

	if err = godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using environment variables or defaults")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode,
	)

	databaseSession, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return &config, err
	}

	databaseSession.AutoMigrate(&dbmodel.Product{})
	config.ProductEntryRepository = dbmodel.NewProductRepository(databaseSession)
	return &config, nil
}
