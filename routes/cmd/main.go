package main

import (
	"log"

	"github.com/Solutions-Corp/chetapp-backend/routes/internal/config"
	"github.com/Solutions-Corp/chetapp-backend/routes/internal/handler"
	"github.com/Solutions-Corp/chetapp-backend/routes/internal/middleware"
	"github.com/Solutions-Corp/chetapp-backend/routes/internal/model"
	"github.com/Solutions-Corp/chetapp-backend/routes/internal/repository"
	"github.com/Solutions-Corp/chetapp-backend/routes/internal/service"
	"github.com/Solutions-Corp/chetapp-backend/routes/internal/websocket"
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
	if err := db.AutoMigrate(&model.Route{}); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	log.Println("Database migration completed")

	routeRepo := repository.NewRouteRepository(db)
	routeService := service.NewRouteService(routeRepo)
	routeHandler := handler.NewRouteHandler(routeService)

	authMiddleware := middleware.AuthMiddleware(&config)

	api := r.Group("/api")
	api.Use(authMiddleware)
	{
		api.POST("/routes", routeHandler.CreateRoute)
		api.GET("/routes", routeHandler.GetAllRoutes)
		api.GET("/routes/:id", routeHandler.GetRouteByID)
		api.POST("/routes/upload-gpx", routeHandler.UploadGPX)
	}

	hub := websocket.NewHub()
	go hub.Run()
	r.GET("/ws", func(c *gin.Context) {
		websocket.ServeWs(hub, c)
	})

	// Inicializa las métricas de Prometheus
	handler.SetupPrometheusMetrics()

	// Endpoints de salud y métricas sin autenticación
	r.GET("/api/health", handler.HealthCheckHandler)
	r.GET("/metrics", handler.PrometheusHandler())

	log.Println("Routes service running on :" + config.Port)
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
