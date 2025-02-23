package services

import (
	"errors"
	"go-demo-app/internal/utils/logger"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(username, password string) error
	Login(username, password string) (string, error)
}

type authService struct {
	// userRepo repositories.UserRepository
}

// NewAuthService creates instance of AuthService

func NewAuthService() AuthService {
	return &authService{
		// userRepo: repo,
	}
}

func (s *authService) Register(username, password string) error {
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	logger.Info.Printf("u: %s hash: %s", username, hash)
	// Save user to database
	return nil
}

func (s *authService) Login(username, password string) (string, error) {
	// Fetch user from DB to get hashed password

	dummyHash, _ := bcrypt.GenerateFromPassword([]byte("secret"), bcrypt.DefaultCost)

	// compare password
	if err := bcrypt.CompareHashAndPassword(dummyHash, []byte(password)); err != nil {
		return "", errors.New("Invalid credentials")
	}

	token, err := CreateJWTToken(username)

	if err != nil {
		return "", err
	}

	return token, nil
}
