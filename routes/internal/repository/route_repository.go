package repository

import (
	"github.com/Solutions-Corp/chetapp-backend/routes/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RouteRepository define la interfaz para operaciones con rutas
type RouteRepository interface {
	CreateRoute(route *model.Route) error
	GetRouteByID(id uuid.UUID) (*model.Route, error)
	GetAllRoutes() ([]model.Route, error)
	UpdateRoute(route *model.Route) error
	DeleteRoute(id uuid.UUID) error
}

// routeRepository implementa la interfaz RouteRepository
type routeRepository struct {
	db *gorm.DB
}

// NewRouteRepository crea una nueva instancia del repositorio de rutas
func NewRouteRepository(db *gorm.DB) RouteRepository {
	return &routeRepository{
		db: db,
	}
}

// CreateRoute crea una nueva ruta en la base de datos
func (r *routeRepository) CreateRoute(route *model.Route) error {
	return r.db.Create(route).Error
}

// GetRouteByID obtiene una ruta por su ID
func (r *routeRepository) GetRouteByID(id uuid.UUID) (*model.Route, error) {
	var route model.Route
	if err := r.db.Where("id = ?", id).First(&route).Error; err != nil {
		return nil, err
	}
	return &route, nil
}

// GetAllRoutes obtiene todas las rutas
func (r *routeRepository) GetAllRoutes() ([]model.Route, error) {
	var routes []model.Route
	if err := r.db.Find(&routes).Error; err != nil {
		return nil, err
	}
	return routes, nil
}

// UpdateRoute actualiza una ruta existente
func (r *routeRepository) UpdateRoute(route *model.Route) error {
	return r.db.Save(route).Error
}

// DeleteRoute elimina una ruta por su ID
func (r *routeRepository) DeleteRoute(id uuid.UUID) error {
	return r.db.Delete(&model.Route{}, id).Error
}
