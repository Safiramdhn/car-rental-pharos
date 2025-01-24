package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string         `json:"name" gorm:"not null;size:255" binding:"required"`
	NIK         string         `json:"nik" gorm:"index:,unique,composite:nik_deleted_at;size:16" binding:"required;maxSize:16"`
	PhoneNumber string         `json:"phone_number" gorm:"not null" binding:"required"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index:,composite:nik_deleted_at" json:"-"`
}