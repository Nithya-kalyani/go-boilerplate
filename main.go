package main

import (
	"log"
	"net/http"

	_ "github.com/Nithya-kalyani/go-boilerplate/docs" // Swagger docs import
	"github.com/Nithya-kalyani/go-boilerplate/internal/cache"
	"github.com/Nithya-kalyani/go-boilerplate/internal/routes"
	services "github.com/Nithya-kalyani/go-boilerplate/internal/services/user-service"
	"github.com/Nithya-kalyani/go-boilerplate/pkg/config"
	"github.com/Nithya-kalyani/go-boilerplate/pkg/db"
	"github.com/Nithya-kalyani/go-boilerplate/pkg/logger"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// Initialize logging
	logger.InitLogger()

	// Load configurations
	cfg := config.LoadConfig()

	// Connect to the database and assign it to the service package's DB variable
	database := db.ConnectDB(cfg)
	services.DB = database

	// Connect to Redis
	cache.ConnectRedis(cfg)

	// Register routes and enable Swagger
	router := routes.RegisterRoutes()
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Start the server
	logger.Log.Info("Server starting...")
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
