package model

import "time"

type User struct {
	ID                 string    `gorm:"primaryKey" json:"id"`
	CorporateAccountNo string    `gorm:"not null;unique" json:"corporate_account_no"`
	CorporateName      string    `gorm:"not null" json:"corporate_name"`
	Name               string    `gorm:"not null" json:"name"`
	Role               string    `gorm:"not null" json:"role"`
	Phone              string    `gorm:"not null" json:"phone"`
	Email              string    `gorm:"not null;unique" json:"email"`
	Password           string    `gorm:"not null" json:"-"`
	Verified           bool      `gorm:"default:false" json:"verified"`
	CreatedAt          time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
