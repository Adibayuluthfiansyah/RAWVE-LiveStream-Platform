package main

import (
	"log"
	"os"

	_ "github.com/Adibayuluthfiansyah/Go-LiveChat/docs"
	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/config"
	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/delivery/http/middleware"
	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/delivery/websocket"
	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/handlers"
	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/repository/postgres"
	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           RAWVE Livestream API
// @version         1.0
// @description     REST API untuk platform livestream RAWVE dengan fitur live chat dan stream management.
// @termsOfService  http://swagger.io/terms/

// @contact.name   RAWVE API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@rawve.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	config.LoadConfig()
	config.ConnectDatabase()

	router := gin.Default()

	//cors
	configCors := cors.DefaultConfig()
	configCors.AllowAllOrigins = true
	configCors.AllowCredentials = true
	configCors.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}
	configCors.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(configCors))

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

	// rate limiter
	api.Use(middleware.RateLimitMiddleware())

	// public route
	// Ping godoc
	// @Summary      Health check
	// @Description  Checks if the server is running
	// @Tags         system
	// @Produce      json
	// @Success      200  {object}  map[string]interface{}  "message"
	// @Router       /ping [get]
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

	// Dashboard godoc
	// @Summary      Creator dashboard
	// @Description  Access to creator studio dashboard for authenticated users
	// @Tags         users
	// @Produce      json
	// @Success      200  {object}  map[string]interface{}  "message, user_id, status"
	// @Failure      401  {object}  map[string]interface{}  "error: Unauthorized"
	// @Security     BearerAuth
	// @Router       /dashboard [get]
	protected.GET("/dashboard", func(c *gin.Context) {
		user_id, _ := c.Get("user_id")
		c.JSON(200, gin.H{
			"message": "Welcome to Creator Studio RAWVE!",
			"user_id": user_id,
			"status":  "success",
		})
	})

	// Swagger documentation route (public)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running at http://localhost:%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
