package model

import "time"

type TransferRecord struct {
	ID                     int          `gorm:"primaryKey;type:int;autoIncrement;" json:"id"`
	TransactionReferenceNo string       `gorm:"not null" json:"transaction_reference_no"`
	TransactionDetail      *Transaction `gorm:"foreignKey:TransactionReferenceNo;references:ReferenceNo" json:"transaction_detail,omitempty"`
	No                     int          `gorm:"not null" json:"no"`
	ToAccountNo            string       `gorm:"not null" json:"to_account_no"`
	ToAccountName          string       `gorm:"not null" json:"to_account_name"`
	ToBankName             string       `gorm:"not null" json:"to_bank_name"`
	Amount                 float64      `gorm:"default:0" json:"transfer_amount"`
	Description            string       `json:"description"`
	Status                 string       `gorm:"default:'AWAITING_APPROVAL'" json:"status"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
