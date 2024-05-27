package model

import "time"

type UserOTP struct {
	ID        int       `gorm:"primaryKey;type:int;autoIncrement;" json:"id"`
	Email     string    `gorm:"not null" json:"email"`
	OTP       string    `gorm:"not null" json:"otp"`
	ExpiredAt time.Time `gorm:"not null" json:"expired_at"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
