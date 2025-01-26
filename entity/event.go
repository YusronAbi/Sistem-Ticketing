package entity

import (
	"time"
)

type Event struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
	Date        time.Time `gorm:"not null"`
	Location    string    `gorm:"type:varchar(255);not null"`
	Price       float64   `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
