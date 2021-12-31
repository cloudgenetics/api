package cloudgenetics

import (
	"time"
)

// User DB User table
type User struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Email       *string
	Userid      string `gorm:"index"`
	EmailVerify bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
