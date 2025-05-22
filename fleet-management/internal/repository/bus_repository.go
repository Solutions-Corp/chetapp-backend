package repository

import (
	"github.com/Solutions-Corp/chetapp-backend/fleet-management/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BusRepository interface {
	CreateBus(bus *model.Bus) error
	GetBusByID(id uuid.UUID) (*model.Bus, error)
	GetAllBuses() ([]model.Bus, error)
}

type busRepository struct {
	db *gorm.DB
}

func NewBusRepository(db *gorm.DB) BusRepository {
	return &busRepository{
		db: db,
	}
}

func (r *busRepository) CreateBus(bus *model.Bus) error {
	return r.db.Create(bus).Error
}

func (r *busRepository) GetBusByID(id uuid.UUID) (*model.Bus, error) {
	var bus model.Bus
	if err := r.db.Preload("Company").Preload("Gps").Where("id = ?", id).First(&bus).Error; err != nil {
		return nil, err
	}
	return &bus, nil
}

func (r *busRepository) GetAllBuses() ([]model.Bus, error) {
	var buses []model.Bus
	if err := r.db.Preload("Company").Preload("Gps").Find(&buses).Error; err != nil {
		return nil, err
	}
	return buses, nil
}
