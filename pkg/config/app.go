package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sai435603/food-app-go/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbInstance *gorm.DB

func connectDb() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Env file can't load at this moment...")
	}

	user := os.Getenv("DB_USER")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	sslmode := os.Getenv("DO_SSLMODE")
	connectionString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		user,
		dbPassword,
		host,
		port,
		dbName,
		sslmode,
	)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to the postgres db successfully done!")
	log.Println("Running database migrations...")
	err = db.AutoMigrate(
		&models.User{},
		&models.Address{},
		&models.Restaurant{},
		&models.Menu{},
		&models.DeliveryPartner{},
		&models.Order{},
		&models.OrderItem{},
		&models.Payment{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migration completed successfully!")
	DbInstance = db

}

func GetDb() *gorm.DB {
	if DbInstance == nil {
		connectDb()
	}
	return DbInstance
}
