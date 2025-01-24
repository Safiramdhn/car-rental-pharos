package models

import "time"

type Membership struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"not null"`
	Discount  int       `json:"discount" gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type MembershipDTO struct {
	ID           uint
	MembershipID uint `json:"membership_id"`
}
