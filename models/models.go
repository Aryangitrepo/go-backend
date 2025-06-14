package models

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// future updates
type DoctorAppointment struct {
	ID          int `gorm:"primaryKey;autoIncrement:true"`
	DoctorId    int
	PatientId   int
	User        User        `gorm:"foreignKey:DoctorId;references:ID"`
	PatientData PatientData `gorm:"foreignKey:PatientId;references:ID"`
}

// for login data
type LoginJson struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Config initializes the database connection
func Config() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("loading .env file: %w", err)
	}

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		return errors.New("DATABASE_URL environment variable not set")
	}

	var err error
	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return fmt.Errorf("connecting to database: %w", err)
	}

	// Migrate all models
	err = db.AutoMigrate(&User{}, &PatientData{}, &DoctorAppointment{})
	if err != nil {
		return fmt.Errorf("migrating database: %w", err)
	}

	log.Println("Database connection established and models migrated")
	return nil
}
