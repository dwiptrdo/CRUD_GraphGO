package database

import (
	"fmt"
	model "graph-go/model"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func InitDB() {
	var err error
	err = godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbName, dbPassword)

	fmt.Println("DB_URI:", dbURI)

	DB, err = gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Println("Error:", err)
		panic("failed to connect to database")
	}

	DB = DB.Debug()

	DB.AutoMigrate(&model.User{})
}
