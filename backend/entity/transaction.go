package entity

import (
	"time"

	"github.com/aliirsyaadn/neobank-technical-test/model"
)

type GetTransactionResponse struct {
	model.Transaction
	TransferRecords []model.TransferRecord `json:"transfer_records"`
}

type GetListTransactionRequest struct {
	Pagination
	Status        string `form:"status"`
	FromAccountNo string `form:"from_account_no"`
	EndDateStr    string `form:"end_date"`
	EndDate       time.Time
	StartDateStr  string `form:"start_date"`
	StartDate     time.Time
}

type GetListTransactionResponse struct {
	Transactions []model.Transaction `json:"transactions"`
	Pagination
}

type CreateTransactionRequest struct {
	MakerUserID         string    // get from jwt
	TotalTransferAmount float64   `json:"total_transfer_amount" binding:"required"`
	TotalTranferRecord  int       `json:"total_transfer_record" binding:"required"`
	TransferDate        time.Time `json:"transfer_date"`
	InstructionType     string    `json:"instruction_type" binding:"required"`

	TransferRecords []model.TransferRecord `json:"transfer_records" binding:"required"`
}

type UpdateTransactionRequest struct {
	ReferenceNo string
	Status      string `json:"status" binding:"required"`
}
