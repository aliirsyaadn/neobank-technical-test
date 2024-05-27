package model

import "time"

type Transaction struct {
	ReferenceNo         string    `gorm:"unique;primaryKey" json:"reference_no"`
	TotalTransferAmount float64   `gorm:"not null" json:"total_transfer_amount"`
	TotalTranferRecord  int       `gorm:"not null" json:"total_transfer_record"`
	FromAccountNo       string    `gorm:"not null" json:"from_account_no"`
	MakerUserID         string    `gorm:"not null" json:"maker_user_id"`
	MakerUserDetail     *User     `gorm:"foreignKey:MakerUserID;references:ID" json:"maker_user_detail,omitempty"`
	TransferDate        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"transfer_date"`
	Status              string    `gorm:"default:'AWAITING_APPROVAL'" json:"status"`
	InstructionType     string    `gorm:"not null" json:"instruction_type"`
	TransferType        string    `gorm:"default:'ONLINE'"  json:"transfer_type"` // additional field : always "Online"
	EstimatedServiceFee float64   `gorm:"default:0" json:"estimated_service_fee"` // additional field : always 0
	CreatedAt           time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
