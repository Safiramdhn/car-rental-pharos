package models

type DriverIncentive struct {
	ID        uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	BookingID uint    `json:"booking_id" gorm:"not null"`
	Incentive float64 `json:"incentive" gorm:"type:numeric;not null"`
}
