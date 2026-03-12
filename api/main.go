package main

import (
	"log"
	"os"

	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/config"
	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/delivery/http/middleware"
	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/handlers"
	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/repository/postgres"
	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.ConnectDatabase()
	router := gin.Default()
	api := router.Group("/api")

	// protected
	protected := api.Group("/")
	protected.Use(middleware.RequireAuth())

	router.GET("/dashboard", func(c *gin.Context) {
		user_id, _ := c.Get("user_id")
		c.JSON(200, gin.H{"message": "This Dashboard Page",
			"user_id": user_id,
			"status":  "success",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Server is running"})
	})

	userRepo := postgres.NewUserRepository(config.DB)
	userUseCase := usecase.NewUserUsecase(userRepo)

	handlers.NewUserHandler(api, userUseCase)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running at http://localhost:%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
