package entity

import (
	"time"
)

type Ticket struct {
	ID        uint    `gorm:"primaryKey"`
	EventID   uint    `gorm:"not null"`
	UserID    uint    `gorm:"not null"`
	Status    string  `gorm:"type:enum('pending','sold','cancelled');default:'pending'"`
	Price     float64 `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Event     Event `gorm:"foreignKey:EventID"`
}
