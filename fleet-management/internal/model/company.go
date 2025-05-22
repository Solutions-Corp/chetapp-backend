package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedBy uuid.UUID `gorm:"type:uuid;not null" json:"created_by"`
	UpdatedBy uuid.UUID `gorm:"type:uuid;not null" json:"updated_by"`
}

func (company *Company) BeforeCreate(tx *gorm.DB) error {
	if company.ID == uuid.Nil {
		company.ID = uuid.New()
	}
	return nil
}
