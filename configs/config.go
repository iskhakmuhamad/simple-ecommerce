package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/iskhakmuhamad/ecommerce/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDatabaseConnection() *gorm.DB {

	err := godotenv.Load("../configs/.env")
	if err != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{LogLevel: logger.Info},
		),
	})

	if err != nil {
		panic("failed to create connection to database")
	}

	db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Cart{},
		&models.Payment{},
	)

	return db
}
