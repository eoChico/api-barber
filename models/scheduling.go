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
	BarberID    uint      `gorm:"not null"`
	Barber      Barber    `gorm:"not null"`
	ClientID    uint      `gorm:"not null"`
	Client      Client    `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
type SchedulingResponse struct {
	ID          uint      `json:"id"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"starttime"`
	EndTime     time.Time `json:"endtime"`
	Status      string    `json:"status"`
	BarberID    uint      `jason:"barber-id"`
	Barber      Barber    `json:"barber"`
	ClientID    uint      `json:"client-id"`
	Client      Client    `json:"client"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
