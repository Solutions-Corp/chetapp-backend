package service

import (
	"encoding/json"
	"fmt"
	"io"

	"strings"
	"time"

	"github.com/Solutions-Corp/chetapp-backend/routes/internal/model"
	"github.com/Solutions-Corp/chetapp-backend/routes/internal/repository"
	"github.com/google/uuid"
	"github.com/tkrajina/gpxgo/gpx"
	"gorm.io/datatypes"
)

type RouteService interface {
	CreateRoute(route *model.Route) error
	GetRouteByID(id uuid.UUID) (*model.Route, error)
	GetAllRoutes() ([]model.Route, error)
	UpdateRoute(route *model.Route) error
	DeleteRoute(id uuid.UUID) error
	ProcessGPXFile(reader io.Reader, name string, createdBy uuid.UUID, updatedBy uuid.UUID) (*model.Route, error)
}

type routeService struct {
	routeRepository repository.RouteRepository
}

func NewRouteService(routeRepository repository.RouteRepository) RouteService {
	return &routeService{
		routeRepository: routeRepository,
	}
}

func (s *routeService) CreateRoute(route *model.Route) error {
	return s.routeRepository.CreateRoute(route)
}

func (s *routeService) GetRouteByID(id uuid.UUID) (*model.Route, error) {
	return s.routeRepository.GetRouteByID(id)
}

func (s *routeService) GetAllRoutes() ([]model.Route, error) {
	return s.routeRepository.GetAllRoutes()
}
func (s *routeService) UpdateRoute(route *model.Route) error {
	return s.routeRepository.UpdateRoute(route)
}

func (s *routeService) DeleteRoute(id uuid.UUID) error {
	return s.routeRepository.DeleteRoute(id)
}

func (s *routeService) ProcessGPXFile(reader io.Reader, name string, createdBy uuid.UUID, updatedBy uuid.UUID) (*model.Route, error) {
	gpxData, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	gpxFile, err := gpx.ParseBytes(gpxData)
	if err != nil {
		return nil, err
	}

	var coordinates []model.Coordinate
	for _, track := range gpxFile.Tracks {
		for _, segment := range track.Segments {
			for _, point := range segment.Points {
				coordinates = append(coordinates, model.Coordinate{
					Latitude:  point.Latitude,
					Longitude: point.Longitude,
				})
			}
		}
	}

	origin := ""
	destination := ""
	if len(coordinates) > 0 {
		origin = formatCoordinate(coordinates[0])
		destination = formatCoordinate(coordinates[len(coordinates)-1])
	}

	coordinatesJSON, err := json.Marshal(coordinates)
	if err != nil {
		return nil, err
	}

	route := &model.Route{
		ID:          uuid.New(),
		Name:        name,
		Origin:      origin,
		Destination: destination,
		Coordinates: datatypes.JSON(coordinatesJSON),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CreatedBy:   createdBy,
		UpdatedBy:   updatedBy,
	}

	if err := s.routeRepository.CreateRoute(route); err != nil {
		return nil, err
	}

	return route, nil
}

func formatCoordinate(coord model.Coordinate) string {
	return strings.TrimSpace(
		fmt.Sprintf("%.6f, %.6f", coord.Latitude, coord.Longitude),
	)
}
