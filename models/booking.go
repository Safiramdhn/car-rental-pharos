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
	ID              uint           `json:"id" gorm:"primaryKey;unique;not null"`
	CustomerID      uint           `json:"customer_id" gorm:"not null" binding:"required"`
	CarID           uint           `json:"car_id" gorm:"not null" binding:"required"`
	BookingTypeID   *uint          `json:"booking_type_id" gorm:"index"` // Optional
	DriverID        *uint          `json:"driver_id" gorm:"index"`       // Optional
	StartRent       time.Time      `json:"start_rent" gorm:"type:date;not null" binding:"required"`
	EndRent         time.Time      `json:"end_rent" gorm:"type:date;not null" binding:"required,gtfield=StartRent"`
	TotalCost       float64        `json:"total_cost" gorm:"type:numeric"`
	TotalDriverCost *float64       `json:"total_driver_cost" gorm:"type:numeric;default:0"` // Optional
	Discount        *float64       `json:"discount" gorm:"type:numeric"`                    // Optional
	IsFinished      bool           `json:"is_finished" gorm:"default:false"`
	CreatedAt       time.Time      `gorm:"autoCreateTime;<-:create" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
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

	// check customer membership discount
	var customer Customer
	if err := tx.Preload("Membership").First(&customer, b.CustomerID).Error; err != nil {
		log.Printf("Failed to get customer with ID %d: %v", b.CustomerID, err)
		return err
	}
	if customer.Membership != nil && customer.Membership.Discount > 0 {
		discount := b.TotalCost * (float64(customer.Membership.Discount) / 100)
		b.Discount = &discount
	} else {
		zero := 0.0
		b.Discount = &zero
	}

	// set total driver cost based on start and end rent
	var driver Driver
	if b.DriverID != nil && *b.DriverID > 0 {
		if err := tx.First(&driver, *b.DriverID).Error; err != nil {
			log.Printf("Failed to get driver with ID %d: %v", *b.DriverID, err)
			return err
		}
		totalDriverCost := driver.DailyCost * b.EndRent.Sub(b.StartRent).Hours() / 24
		b.TotalDriverCost = &totalDriverCost
		if *b.TotalDriverCost <= 0 {
			*b.TotalDriverCost = driver.DailyCost
		}
	} else {
		zero := 0.0
		b.TotalDriverCost = &zero
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
		log.Println("true")
		b.IsFinished = true
	} else {
		log.Println("false")
		b.IsFinished = false
	}

	// check customer membership discount
	var customer Customer
	if err := tx.Preload("Membership").First(&customer, b.CustomerID).Error; err != nil {
		log.Printf("Failed to get customer with ID %d: %v", b.CustomerID, err)
		return err
	}
	if customer.Membership != nil && customer.Membership.Discount > 0 {
		discount := b.TotalCost * (float64(customer.Membership.Discount) / 100)
		b.Discount = &discount
	} else {
		zero := 0.0
		b.Discount = &zero
	}

	var driver Driver
	if b.DriverID != nil && *b.DriverID > 0 {
		if err := tx.First(&driver, *b.DriverID).Error; err != nil {
			log.Printf("Failed to get driver with ID %d: %v", *b.DriverID, err)
			return err
		}
		totalDriverCost := driver.DailyCost * b.EndRent.Sub(b.StartRent).Hours() / 24
		b.TotalDriverCost = &totalDriverCost
		if *b.TotalDriverCost <= 0 {
			*b.TotalDriverCost = driver.DailyCost
		}
	} else {
		zero := 0.0
		b.TotalDriverCost = &zero
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

		// add driver incentive record
		if b.DriverID != nil && *b.DriverID > 0 {
			countIncentive := b.TotalCost * (5.0 / 100.0)

			// if booking id already exists
			var existedRecord DriverIncentive
			err := tx.First(&existedRecord, b.ID).Error
			if err != nil {

				record := DriverIncentive{
					BookingID: b.ID,
					Incentive: countIncentive,
				}
				if err := tx.Create(&record).Error; err != nil {
					log.Printf("Failed to create driver incentive record for booking ID %d: %v", b.ID, err)
					return err
				}
			} else {
				// update existing record
				if err := tx.Model(DriverIncentive{}).Where("booking_id = ?", b.ID).Set("incentive = incentive +?", countIncentive).Update("incentive", gorm.Expr("incentive + ?", countIncentive)).Error; err != nil {
					log.Printf("Failed to update driver incentive record for booking ID %d: %v", b.ID, err)
					return err
				}
			}
		}
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
