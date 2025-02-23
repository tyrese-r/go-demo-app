package db

import (
	"fmt"
	"go-demo-app/internal/utils/logger"
	"log"
	"sync"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// Database struct to hold the GORM instance
type Database struct {
	DB *gorm.DB
}

var (
	database *Database
	once     sync.Once
)

// GetDB returns the singleton GORM database instance
func GetDB() *Database {
	return database
}

// ConnectToDatabase initializes and connects to the SQLite database
func ConnectToDatabase() (*Database, error) {
	var err error

	once.Do(func() {
		db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
		if err != nil {
			logger.Error.Fatalf("Failed to open db: %v", err)
		}

		database = &Database{DB: db}
		log.Println("Connected to SQLite database!")
	})

	return database, err
}

// CloseDatabase closes the database connection (GORM handles this automatically, but included for completeness)
func CloseDatabase() {
	sqlDB, err := database.DB.DB()
	if err != nil {
		logger.Error.Fatalf("Error getting SQL DB instance: %v", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		logger.Error.Fatalf("Error closing database: %v", err)
	}

	fmt.Println("Database connection closed.")
}
