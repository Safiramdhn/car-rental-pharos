package models

import "time"

type BookingType struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Type        string    `json:"type" gorm:"not null"`
	Description string    `json:"description"` // Optional
	CreatedAt   time.Time `gorm:"autoCreateTime;<-:create" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
