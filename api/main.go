package main

import (
	"log"
	"os"

	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/config"
	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/delivery/http/middleware"
	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/delivery/websocket"
	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/handlers"
	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/repository/postgres"
	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.ConnectDatabase()

	router := gin.Default()

	// inject
	userRepo := postgres.NewUserRepository(config.DB)
	userUseCase := usecase.NewUserUsecase(userRepo)

	chatRepo := postgres.NewChatRepository(config.DB)
	chatUseCase := usecase.NewChatUsecase(chatRepo)

	// websocket
	hub := websocket.NewHub(chatUseCase)
	go hub.Run()

	// route
	api := router.Group("/api")

	// public route
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Server RAWVE is running!"})
	})

	streamHandler := handlers.StreamHandler{ChatUsecase: chatUseCase}
	api.GET("/streams/live", streamHandler.GetLiveStream)

	handlers.NewUserHandler(api, userUseCase)

	api.GET("/ws/chat/:stream_id", func(c *gin.Context) {
		websocket.ServeWS(hub, c)
	})

	// route clerk
	protected := api.Group("/")
	protected.Use(middleware.RequireAuth())

	// handler protected route
	handlers.NewStreamHandler(protected, chatUseCase)

	protected.PUT("/profile/setup", func(c *gin.Context) {
		userHandler := handlers.UserHandler{UserUsecase: userUseCase}
		userHandler.SetupProfile(c)
	})

	protected.GET("/dashboard", func(c *gin.Context) {
		user_id, _ := c.Get("user_id")
		c.JSON(200, gin.H{
			"message": "Welcome to Creator Studio RAWVE!",
			"user_id": user_id,
			"status":  "success",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running at http://localhost:%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
