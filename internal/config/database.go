package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	err = database.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatal("Failed migration database:", err)
	}

	fmt.Println("Connection and Database success")
	DB = database
}
