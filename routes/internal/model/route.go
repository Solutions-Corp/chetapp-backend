package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Coordinate struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

type Route struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string         `gorm:"type:varchar(255);not null" json:"name"`
	Origin      string         `gorm:"type:varchar(255);not null" json:"origin"`
	Destination string         `gorm:"type:varchar(255);not null" json:"destination"`
	Coordinates datatypes.JSON `gorm:"type:jsonb;not null" json:"coordinates"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedBy   uuid.UUID      `gorm:"type:uuid;not null" json:"created_by"`
	UpdatedBy   uuid.UUID      `gorm:"type:uuid;not null" json:"updated_by"`
}

func (route *Route) BeforeCreate(tx *gorm.DB) error {
	if route.ID == uuid.Nil {
		route.ID = uuid.New()
	}
	return nil
}
