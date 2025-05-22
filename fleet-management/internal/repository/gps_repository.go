package repository

import (
	"github.com/Solutions-Corp/chetapp-backend/fleet-management/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GpsRepository interface {
	CreateGps(gps *model.Gps) error
	GetGpsByID(id uuid.UUID) (*model.Gps, error)
	GetAllGps() ([]model.Gps, error)
}

type gpsRepository struct {
	db *gorm.DB
}

func NewGpsRepository(db *gorm.DB) GpsRepository {
	return &gpsRepository{
		db: db,
	}
}

func (r *gpsRepository) CreateGps(gps *model.Gps) error {
	return r.db.Create(gps).Error
}

func (r *gpsRepository) GetGpsByID(id uuid.UUID) (*model.Gps, error) {
	var gps model.Gps
	if err := r.db.Where("id = ?", id).First(&gps).Error; err != nil {
		return nil, err
	}
	return &gps, nil
}

func (r *gpsRepository) GetAllGps() ([]model.Gps, error) {
	var gpss []model.Gps
	if err := r.db.Find(&gpss).Error; err != nil {
		return nil, err
	}
	return gpss, nil
}
