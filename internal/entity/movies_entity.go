package entity

import (
	"time"
)

type Movies struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"type:varchar(100)" validate:"required"`
	Description string    `json:"description" gorm:"type:varchar(255)" validate:"required"`
	Rating      float64   `json:"rating" validate:"required"`
	Image       string    `json:"image"`
	Created_at  time.Time `json:"createdAt"`
	Updated_at  time.Time `json:"updatedAt"`
}

func (Movies) TableName() string { return "movie" }
