package main

import (
	"github.com/gin-gonic/gin"
	"go-demo-app/internal/db"
	"go-demo-app/internal/db/schema"
	"go-demo-app/internal/handlers"
	"go-demo-app/internal/services"
	"go-demo-app/internal/utils/logger"
	"go-demo-app/internal/utils/secrets"
)

func main() {
	logger.InitLogger("app.log")
	secrets.LoadEnv()
	router := gin.Default()

	// Create db
	_, err := db.ConnectToDatabase()

	schema.MigrateUserTable()

	if err != nil {
		logger.Error.Fatalf("Failed to connect to database: %v", err)
	}

	defer db.CloseDatabase()

	authService := services.NewAuthService()
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler()

	api := router.Group("/api")
	{
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)

		api.GET("/users/:username", userHandler.GetUserHandler)
		api.POST("/users", userHandler.CreateUserHandler)
	}

	logger.Info.Println("Hello world!")

	err = router.Run(":8080")
	if err != nil {
		return
	}

}
