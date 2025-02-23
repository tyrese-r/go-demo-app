package schema

import (
	"go-demo-app/internal/db"
	"go-demo-app/internal/utils/logger"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"log"
)

// User represents the users table
type User struct {
	gorm.Model
	Username string         `json:"username" gorm:"uniqueIndex;not null"`
	Password string         `json:"password" gorm:"not null"`
	Roles    datatypes.JSON `json:"roles" gorm:"type:json"`
}

// MigrateUserTable ensures the `users` table exists
func MigrateUserTable() {
	dbInstance := db.GetDB()
	err := dbInstance.DB.AutoMigrate(&User{})
	if err != nil {
		logger.Error.Fatalf("Error migrating users table: %v", err)
	}
	log.Println("Users table migrated successfully!")
}
