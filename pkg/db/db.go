package db

import (
	"fmt"
	"log"

	"github.com/Nithya-kalyani/go-boilerplate/internal/models"
	"github.com/Nithya-kalyani/go-boilerplate/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	DB = db

	// Auto migrate models
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to auto-migrate model: %v", err)
	}
	log.Println("Database connection established successfully.")
	return db
}
