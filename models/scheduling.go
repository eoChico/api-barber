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
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
type SchedulingResponse struct {
	ID          uint      `json:"id"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"starttime"`
	EndTime     time.Time `json:"endtime"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
