package models

import "time"

type Client struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"not null;unique"`
	Phone     string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ClientResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	phone string `json:"phone"`
}
