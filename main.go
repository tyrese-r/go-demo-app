package main

import (
	"github.com/gin-gonic/gin"
	"go-demo-app/internal/handlers"
	"go-demo-app/internal/services"
	"go-demo-app/internal/utils/logger"
	"go-demo-app/internal/utils/secrets"
)

func main() {
	logger.InitLogger("app.log")
	secrets.LoadEnv()
	router := gin.Default()

	authService := services.NewAuthService()
	authHandler := handlers.NewAuthHandler(authService)

	api := router.Group("/api")
	{
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)
	}

	logger.Info.Println("Hello world!")

	err := router.Run(":8080")
	if err != nil {
		return
	}

}
