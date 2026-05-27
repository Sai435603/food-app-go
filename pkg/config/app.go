package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbInstance *gorm.DB

func connectDb() {
	err := godotenv.Load("../ ../.env")
	if err != nil {
		log.Fatal("Env file can't load at this moment...")
	}

	user := os.Getenv("DB_USER")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	connectionString := fmt.Sprintf("user=%s host=%s port=%s name=%s password=%s", user, host, port, dbName, dbPassword)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to the postgres db successfully done!")

	log.Fatal("Database Created")
	DbInstance = db
}

func GetDb() *gorm.DB {
	if DbInstance == nil {
		connectDb()
	}
	return DbInstance
}
