package model

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBClient struct {
	DB *gorm.DB
}

var dbClient *DBClient

func InitDatabase() *gorm.DB{
	// import .env variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	// load database url
	postgresURL := os.Getenv("postgresURL")
	// connect to database
	db, err := gorm.Open(postgres.Open(postgresURL), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	// Migrate the schema
	dbClient = &DBClient{
		DB: db,
	}
	return db
}

func GetDB() *DBClient {
	return dbClient
}