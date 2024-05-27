package repository

import (
	"context"

	"github.com/aliirsyaadn/neobank-technical-test/entity"
	"github.com/aliirsyaadn/neobank-technical-test/model"
)

type UserRepository interface {
	GetList(ctx context.Context, req entity.GetListUserRequest) ([]model.User, error)
	GetByEmail(ctx context.Context, email string) (model.User, error)
	Get(ctx context.Context, id string) (model.User, error)
	Insert(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
}

type UserOTPRepository interface {
	Get(ctx context.Context, email, otp string) (model.UserOTP, error)
	Insert(ctx context.Context, otpUser *model.UserOTP) error
}

type TransactionRepository interface {
	Get(ctx context.Context, referenceNo string) (model.Transaction, error)
	GetList(ctx context.Context, req entity.GetListTransactionRequest) (entity.GetListTransactionResponse, error)
	Insert(ctx context.Context, transaction *model.Transaction) error
	Update(ctx context.Context, transaction *model.Transaction) error
}

type TransferRecordRepository interface {
	GetList(ctx context.Context, referenceNo string) ([]model.TransferRecord, error)
	InsertBatch(ctx context.Context, transferRecords *[]model.TransferRecord) error
	UpdateStatus(ctx context.Context, referenceNo string, status string) error
}
