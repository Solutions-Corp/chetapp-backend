package service

import (
	"github.com/Solutions-Corp/chetapp-backend/fleet-management/internal/model"
	"github.com/Solutions-Corp/chetapp-backend/fleet-management/internal/repository"
	"github.com/google/uuid"
)

type BusService interface {
	CreateBus(bus *model.Bus) error
	GetBusByID(id uuid.UUID) (*model.Bus, error)
	GetAllBuses() ([]model.Bus, error)
}

type busService struct {
	busRepo repository.BusRepository
}

func NewBusService(busRepo repository.BusRepository) BusService {
	return &busService{
		busRepo: busRepo,
	}
}

func (s *busService) CreateBus(bus *model.Bus) error {
	return s.busRepo.CreateBus(bus)
}

func (s *busService) GetBusByID(id uuid.UUID) (*model.Bus, error) {
	return s.busRepo.GetBusByID(id)
}

func (s *busService) GetAllBuses() ([]model.Bus, error) {
	return s.busRepo.GetAllBuses()
}
