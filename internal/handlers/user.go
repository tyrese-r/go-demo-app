package handlers

import (
	"github.com/gin-gonic/gin"
	"go-demo-app/internal/db/schema"
	"go-demo-app/internal/repositories"
	"net/http"
)

// UserHandler handles HTTP requests for users
type UserHandler struct {
	UserRepo repositories.UserRepository
}

// NewUserHandler initializes a UserHandler
func NewUserHandler() *UserHandler {
	return &UserHandler{
		UserRepo: repositories.NewUserRepository(),
	}
}

// CreateUserHandler handles user creation via API
func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	var user schema.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.UserRepo.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// GetUserHandler fetches a user by username
func (h *UserHandler) GetUserHandler(c *gin.Context) {
	username := c.Param("username")

	user, err := h.UserRepo.GetUser(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
