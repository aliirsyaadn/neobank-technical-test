package service

import (
	"context"

	"github.com/aliirsyaadn/neobank-technical-test/entity"
	"github.com/aliirsyaadn/neobank-technical-test/model"
)

type UserService interface {
	Login(ctx context.Context, req entity.LoginUserRequest) (res entity.LoginUserResponse, err error)
	Register(ctx context.Context, req entity.RegisterUserRequest) (res entity.RegisterUserResponse, err error)
	SendOTP(ctx context.Context, req entity.SendOTPRequest) error
}

type TransactionService interface {
	Get(ctx context.Context, referenceNo string) (res entity.GetTransactionResponse, err error)
	GetList(ctx context.Context, req entity.GetListTransactionRequest) (res entity.GetListTransactionResponse, err error)
	Create(ctx context.Context, req entity.CreateTransactionRequest) (res model.Transaction, err error)
	Update(ctx context.Context, req entity.UpdateTransactionRequest) (res model.Transaction, err error)
	ValidateTransferRecordCSV(ctx context.Context, records []entity.TransferRecordData) (res entity.ValidateTransferRecordCSVResponse, err error)
}
