package repositories

import (
	"go-demo-app/internal/db"
	"go-demo-app/internal/db/schema"
)

// UserRepository defines database operations for users
type UserRepository interface {
	CreateUser(user *schema.User) error
	GetUser(username string) (*schema.User, error)
}

type userRepo struct {
	db *db.Database
}

// NewUserRepository creates a new UserRepository
func NewUserRepository() UserRepository {
	return &userRepo{
		db: db.GetDB(), // Uses the GORM database instance
	}
}

// CreateUser inserts a new user into the database
func (r *userRepo) CreateUser(user *schema.User) error {
	return r.db.DB.Create(user).Error
}

// GetUser retrieves a user by username
func (r *userRepo) GetUser(username string) (*schema.User, error) {
	var user schema.User
	err := r.db.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
