package models

import (
	"car-rental/helpers"
	"encoding/json"
	"fmt"
	"time"

	"log"

	"gorm.io/gorm"
)

type Booking struct {
	ID         uint           `json:"id" gorm:"primaryKey;unique;not null"`
	CustomerID uint           `json:"customer_id" gorm:"not null" binding:"required"`
	CarID      uint           `json:"car_id" gorm:"not null" binding:"required"`
	StartRent  time.Time      `json:"start_rent" gorm:"type:date;not null" binding:"required"`
	EndRent    time.Time      `json:"end_rent" gorm:"type:date;not null" binding:"required,gtfield=StartRent"`
	TotalCost  float64        `json:"total_cost" gorm:"type:numeric"`
	IsFinished bool           `json:"is_finished" gorm:"default:false"`
	CreatedAt  time.Time      `gorm:"autoCreateTime;<-:create" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

// Custom UnmarshalJSON for StartRent and EndRent
func (b *Booking) UnmarshalJSON(data []byte) error {
	type Alias Booking
	aux := &struct {
		StartRent string `json:"start_rent"`
		EndRent   string `json:"end_rent"`
		*Alias
	}{
		Alias: (*Alias)(b),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	startRent := helpers.FormatDate(aux.StartRent)
	b.StartRent = startRent

	endRent := helpers.FormatDate(aux.EndRent)
	b.EndRent = endRent

	return nil
}

func (b *Booking) BeforeCreate(tx *gorm.DB) (err error) {
	var car Car
	if err := tx.First(&car, b.CarID).Error; err != nil {
		log.Printf("Failed to get car with ID %d: %v", b.CarID, err)
		return err
	}

	// Check if the car is available
	if car.Stock <= 0 {
		return fmt.Errorf("car with ID %d is out of stock", b.CarID)
	}

	// Calculate the total cost
	b.TotalCost = car.DailyRent * b.EndRent.Sub(b.StartRent).Hours() / 24
	if b.TotalCost <= 0 {
		b.TotalCost = car.DailyRent
	}

	// Set IsFinished based on the EndRent date
	if time.Now().After(b.EndRent) {
		b.IsFinished = true
	} else {
		b.IsFinished = false
	}

	return nil
}

func (b *Booking) BeforeUpdate(tx *gorm.DB) (err error) {
	var car Car
	if err := tx.First(&car, b.CarID).Error; err != nil {
		log.Printf("Failed to get car with ID %d: %v", b.CarID, err)
		return err
	}

	// Calculate the total cost
	b.TotalCost = car.DailyRent * b.EndRent.Sub(b.StartRent).Hours() / 24
	if b.TotalCost <= 0 {
		b.TotalCost = car.DailyRent
	}

	// Set IsFinished based on the EndRent date
	if !time.Now().Truncate(24 * time.Hour).Before(b.EndRent.Truncate(24 * time.Hour)) {
		b.IsFinished = true
	} else {
		b.IsFinished = false
	}

	return nil
}

func (b *Booking) AfterSave(tx *gorm.DB) error {
	var car Car
	if err := tx.First(&car, b.CarID).Error; err != nil {
		log.Printf("Failed to get car with ID %d: %v", b.CarID, err)
		return err
	}

	// Adjust the car stock based on the IsFinished status
	if b.IsFinished {
		car.Stock += 1
	} else {
		car.Stock -= 1
	}

	if err := tx.Save(&car).Error; err != nil {
		log.Printf("Failed to update stock for car with ID %d: %v", b.CarID, err)
		return err
	}

	return nil
}

func (b *Booking) AfterDelete(tx *gorm.DB) (err error) {
	var car Car
	if err := tx.First(&car, b.CarID).Error; err != nil {
		log.Printf("Failed to get car with ID %d: %v", b.CarID, err)
		return err
	}

	// Increment the car stock
	car.Stock += 1
	if err := tx.Save(&car).Error; err != nil {
		log.Printf("Failed to update stock for car with ID %d: %v", b.CarID, err)
		return err
	}

	return nil
}
