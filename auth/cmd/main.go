package main

import (
	"log"

	"github.com/Solutions-Corp/chetapp-backend/auth/internal/config"
	"github.com/Solutions-Corp/chetapp-backend/auth/internal/handler"
	"github.com/Solutions-Corp/chetapp-backend/auth/internal/model"
	"github.com/Solutions-Corp/chetapp-backend/auth/internal/repository"
	"github.com/Solutions-Corp/chetapp-backend/auth/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := setupDatabase()

	if err != nil {
		log.Fatalf("Error setting up database: %v", err)
	}

	log.Println("Database connection established")
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	log.Println("Database migration completed")

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	if err := userService.CreateDefaultUser(); err != nil {
		log.Fatalf("Error creating default user: %v", err)
	}

	authHandler := handler.NewAuthHandler(userService, &config)

	r.GET("/health", handler.HealthCheckHandler)
	r.POST("/login", authHandler.Login)

	log.Println("Auth service running on :" + config.Port)
	log.Fatal(r.Run(":" + config.Port))
}

func setupDatabase() (*gorm.DB, error) {
	config, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	dsn := config.DatabaseURL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		return nil, err
	}

	return db, nil
}
