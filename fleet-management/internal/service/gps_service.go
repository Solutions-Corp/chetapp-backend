package service

import (
	"github.com/Solutions-Corp/chetapp-backend/fleet-management/internal/model"
	"github.com/Solutions-Corp/chetapp-backend/fleet-management/internal/repository"
	"github.com/google/uuid"
)

type GpsService interface {
	CreateGps(gps *model.Gps) error
	GetGpsByID(id uuid.UUID) (*model.Gps, error)
	GetAllGps() ([]model.Gps, error)
}

type gpsService struct {
	gpsRepo repository.GpsRepository
}

func NewGpsService(gpsRepo repository.GpsRepository) GpsService {
	return &gpsService{
		gpsRepo: gpsRepo,
	}
}

func (s *gpsService) CreateGps(gps *model.Gps) error {
	return s.gpsRepo.CreateGps(gps)
}

func (s *gpsService) GetGpsByID(id uuid.UUID) (*model.Gps, error) {
	return s.gpsRepo.GetGpsByID(id)
}

func (s *gpsService) GetAllGps() ([]model.Gps, error) {
	return s.gpsRepo.GetAllGps()
}
