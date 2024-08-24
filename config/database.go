package config

import (
	"fmt"
	"go-authentication/models"
	"go-authentication/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		utils.Logger.Printf("Error: %v", err)
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&sql_mode='NO_ZERO_IN_DATE,NO_ZERO_DATE'",dbUser, dbPassword, dbHost, dbPort, dbName)
	// dsn := "root:root@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.Logger.Printf("Error: %v", err)
		log.Fatal("Failed to connect to database", err)
	}
	fmt.Println("Database connected successfully")

	// Auto migrate the models
	DB.AutoMigrate(&models.User{})
}