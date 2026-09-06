package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"type:varchar(255);not null;uniqueIndex"`
	Phone     string    `gorm:"type:varchar(30);"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
