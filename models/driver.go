package models

import (
	"time"

	"gorm.io/gorm"
)

type Driver struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string         `json:"name" gorm:"not null;size:255" binding:"required"`
	NIK         string         `json:"nik" gorm:"index:,unique;size:16" binding:"required"`
	PhoneNumber string         `json:"phone_number" gorm:"not null"`
	DailyCost   float64        `json:"daily_cost" gorm:"type:numeric" binding:"required,gte=0"`
	CreatedAt   time.Time      `gorm:"autoCreateTime;<-:create" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
