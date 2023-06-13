package db

import (
	"fmt"
	"github.com/jedavard/gomotions/pkg/models"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	database, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		panic("Failed to connect to database!")
	}

	// migrate
	log.Println("Perform AutoMigrate")
	database.AutoMigrate(&models.Promotion{})

	DB = database
}
