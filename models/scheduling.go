package models

import (
	"time"
)

type Scheduling struct {
	ID          uint      `gorm:"primaryKey"`
	Description string    `gorm:"not null"`
	StartTime   time.Time `gorm:"not null"`
	EndTime     time.Time `gorm:"not null"`
	Status      string    `gorm:"not null"`
}
