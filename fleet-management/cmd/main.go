package main

import (
	"log"

	"github.com/Solutions-Corp/chetapp-backend/fleet-management/internal/config"
	"github.com/Solutions-Corp/chetapp-backend/fleet-management/internal/handler"
	"github.com/Solutions-Corp/chetapp-backend/fleet-management/internal/middleware"
	"github.com/Solutions-Corp/chetapp-backend/fleet-management/internal/model"
	"github.com/Solutions-Corp/chetapp-backend/fleet-management/internal/repository"
	"github.com/Solutions-Corp/chetapp-backend/fleet-management/internal/service"
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
	if err := db.AutoMigrate(&model.Company{}); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
	if err := db.AutoMigrate(&model.Bus{}); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
	if err := db.AutoMigrate(&model.Gps{}); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	log.Println("Database migration completed")
	companyRepository := repository.NewCompanyRepository(db)
	companyService := service.NewCompanyService(companyRepository)
	companyHandler := handler.NewCompanyHandler(companyService)
	busRepository := repository.NewBusRepository(db)
	busService := service.NewBusService(busRepository)
	busHandler := handler.NewBusHandler(busService)

	gpsRepository := repository.NewGpsRepository(db)
	gpsService := service.NewGpsService(gpsRepository)
	gpsHandler := handler.NewGpsHandler(gpsService)
	// Inicializa las métricas de Prometheus
	handler.SetupPrometheusMetrics()

	// Endpoints de salud y métricas sin autenticación
	r.GET("/api/health", handler.HealthCheckHandler)
	r.GET("/metrics", handler.PrometheusHandler())

	authMiddleware := middleware.AuthMiddleware(&config)

	api := r.Group("/api")
	api.Use(authMiddleware)
	{
		api.POST("/companies", companyHandler.CreateCompany)
		api.GET("/companies/:id", companyHandler.GetCompany)
		api.GET("/companies", companyHandler.GetAllCompanies)
		api.PUT("/companies/:id", companyHandler.UpdateCompany)
		api.DELETE("/companies/:id", companyHandler.DeleteCompany)
		api.POST("/buses", busHandler.CreateBus)
		api.GET("/buses/:id", busHandler.GetBusByID)
		api.GET("/buses", busHandler.GetAllBuses)

		api.POST("/gps", gpsHandler.CreateGps)
		api.GET("/gps/:id", gpsHandler.GetGpsByID)
		api.GET("/gps", gpsHandler.GetAllGps)
	}

	log.Println("Fleet management service running on :" + config.Port)
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
