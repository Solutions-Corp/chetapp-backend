package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Gps struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	RouteID   uuid.UUID `gorm:"type:uuid;not null" json:"route_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedBy uuid.UUID `gorm:"type:uuid;not null" json:"created_by"`
	UpdatedBy uuid.UUID `gorm:"type:uuid;not null" json:"updated_by"`
}

func (gps *Gps) BeforeCreate(tx *gorm.DB) error {
	if gps.ID == uuid.Nil {
		gps.ID = uuid.New()
	}
	return nil
}
