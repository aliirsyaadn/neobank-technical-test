package entity

type ValidateTransferRecordCSVResponse struct {
	TransferRecords     []TransferRecordData `json:"transfer_records,omitempty"`
	TotalTranferRecord  int                  `json:"total_transfer_record"`
	TotalTransferAmount float64              `json:"total_transfer_amount"`
	Errors              []string             `json:"errors,omitempty"`
}

type TransferRecordData struct {
	No            int     `csv:"no" json:"no" binding:"required"`
	ToBankName    string  `csv:"to_bank_name" json:"to_bank_name" binding:"required"`
	ToAccountNo   string  `csv:"to_account_no" json:"to_account_no" binding:"required"`
	ToAccountName string  `csv:"to_account_name" json:"to_account_name" binding:"required"`
	Amount        float64 `csv:"transfer_amount" json:"transfer_amount" binding:"required"`
}
