package models

import (
	"github.com/google/uuid"
	"time"
)

// Promotion example
type Promotion struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	Price     float64   `json:"price"`
	ExpiresAt time.Time `json:"expiresAt" gorm:"type:datetime(6);index"`
}
