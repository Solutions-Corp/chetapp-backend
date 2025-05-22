package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Bus struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Plate     string     `gorm:"type:varchar(255);not null" json:"plate"`
	CompanyID uuid.UUID  `gorm:"type:uuid;not null" json:"company_id"`
	GpsID     *uuid.UUID `gorm:"type:uuid;null" json:"gps_id,omitempty"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedBy uuid.UUID  `gorm:"type:uuid;not null" json:"created_by"`
	UpdatedBy uuid.UUID  `gorm:"type:uuid;not null" json:"updated_by"`
	Company   Company    `gorm:"foreignKey:CompanyID;references:ID" json:"company"`
	Gps       *Gps       `gorm:"foreignKey:GpsID;references:ID" json:"gps,omitempty"`
}

func (bus *Bus) BeforeCreate(tx *gorm.DB) error {
	if bus.ID == uuid.Nil {
		bus.ID = uuid.New()
	}
	return nil
}
