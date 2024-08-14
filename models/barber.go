package models

import "time"

type Barber struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"not null;unique"`
	Phone     string `gorm:"not null"`
	CreatedAt time.Time
	UpdateAt  time.Time
}

type BarberResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
