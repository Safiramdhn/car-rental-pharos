package models

import (
	"time"

	"gorm.io/gorm"
)

type Car struct {
	ID        uint           `json:"id" gorm:"primaryKey;unique;not null"`
	Name      string         `json:"name" gorm:"type:varchar(255);not null" binding:"required"`
	Stock     int            `json:"stock" gorm:"not null;default:0" binding:"required,gte=0"`
	DailyRent float64        `json:"daily_rent" gorm:"type:numeric;not null;default:0" binding:"required,gte=0"`
	CreatedAt time.Time      `gorm:"autoCreateTime;<-:create" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
