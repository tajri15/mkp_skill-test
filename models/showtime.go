package models

import (
	"time"
)

type Showtime struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	MovieID   uint      `json:"movie_id"`
	TheaterID uint      `json:"theater_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Price     float64   `gorm:"type:numeric(10,2)" json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}