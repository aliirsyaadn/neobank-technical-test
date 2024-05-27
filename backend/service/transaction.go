package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/aliirsyaadn/neobank-technical-test/constant"
	"github.com/aliirsyaadn/neobank-technical-test/entity"
	"github.com/aliirsyaadn/neobank-technical-test/model"
	"github.com/aliirsyaadn/neobank-technical-test/repository"
	"go.uber.org/zap"
)

type transactionService struct {
	log                      *zap.SugaredLogger
	userRepository           repository.UserRepository
	transactionRepository    repository.TransactionRepository
	transferRecordRepository repository.TransferRecordRepository
}

func NewTransactionService(
	log *zap.SugaredLogger,
	userRepository repository.UserRepository,
	transactionRepository repository.TransactionRepository,
	transferRecordRepository repository.TransferRecordRepository,
) TransactionService {

	return &transactionService{
		log:                      log,
		userRepository:           userRepository,
		transactionRepository:    transactionRepository,
		transferRecordRepository: transferRecordRepository,
	}
}

func (s *transactionService) Get(ctx context.Context, referenceNo string) (res entity.GetTransactionResponse, err error) {
	transaction, err := s.transactionRepository.Get(ctx, referenceNo)
	if err != nil {
		s.log.Error("Failed to get transaction", err.Error())
		return res, err
	}

	transferRecords, err := s.transferRecordRepository.GetList(ctx, referenceNo)
	if err != nil {
		s.log.Error("Failed to get transfer records", err.Error())
		return res, err
	}

	res = entity.GetTransactionResponse{
		Transaction:     transaction,
		TransferRecords: transferRecords,
	}

	return res, nil
}

func (s *transactionService) GetList(ctx context.Context, req entity.GetListTransactionRequest) (res entity.GetListTransactionResponse, err error) {
	if req.Status != "" {
		_, ok := constant.TRANSFER_RECORD_STATUSES[req.Status]
		if !ok {
			s.log.Error("Invalid status")
			return res, errors.New("invalid status")
		}
	}

	res, err = s.transactionRepository.GetList(ctx, req)
	if err != nil {
		s.log.Error("Failed to get list of transactions", err.Error())
		return res, err
	}

	return res, nil
}

func (s *transactionService) Create(ctx context.Context, req entity.CreateTransactionRequest) (res model.Transaction, err error) {
	valid, ok := constant.TRANSACTION_INSTRUCTION_TYPES[req.InstructionType]
	if !valid || !ok {
		s.log.Error("Invalid instruction type")
		return res, errors.New("invalid instruction type")
	}

	user, err := s.userRepository.Get(ctx, req.MakerUserID)
	if err != nil {
		s.log.Error("Failed to get user", err.Error())
		return res, err
	}

	if req.InstructionType == constant.TRANSACTION_INSTRUCTION_TYPE_IMMEDIATE {
		req.TransferDate = time.Now()
	} else {
		if req.TransferDate.IsZero() {
			s.log.Error("Transfer date is required")
			return res, errors.New("transfer date is required")
		}

	}

	// TODO: create db transaction
	referenceNo := s.generateReferenceNo()
	transaction := model.Transaction{
		ReferenceNo:         referenceNo,
		FromAccountNo:       user.CorporateAccountNo,
		MakerUserID:         req.MakerUserID,
		TotalTransferAmount: req.TotalTransferAmount,
		TotalTranferRecord:  req.TotalTranferRecord,
		InstructionType:     req.InstructionType,
	}

	err = s.transactionRepository.Insert(ctx, &transaction)
	if err != nil {
		s.log.Error("Failed to insert transaction", err.Error())
		return res, err
	}

	for i := range req.TransferRecords {
		if req.TransferRecords[i].No == 0 {
			req.TransferRecords[i].No = i + 1
		}
		req.TransferRecords[i].TransactionReferenceNo = referenceNo
	}

	err = s.transferRecordRepository.InsertBatch(ctx, &req.TransferRecords)
	if err != nil {
		s.log.Error("Failed to insert transfer records", err.Error())
		return res, err
	}

	return transaction, nil
}

func (s *transactionService) Update(ctx context.Context, req entity.UpdateTransactionRequest) (res model.Transaction, err error) {
	valid, ok := constant.TRANSFER_RECORD_STATUSES[req.Status]
	if !valid || !ok {
		s.log.Error("Invalid status")
		return model.Transaction{}, errors.New("invalid status")
	}

	transaction, err := s.transactionRepository.Get(ctx, req.ReferenceNo)
	if err != nil {
		s.log.Error("Failed to get transaction", err.Error())
		return model.Transaction{}, err
	}

	transaction.Status = req.Status

	err = s.transactionRepository.Update(ctx, &transaction)
	if err != nil {
		s.log.Error("Failed to update transaction", err.Error())
		return model.Transaction{}, err
	}

	err = s.transferRecordRepository.UpdateStatus(ctx, req.ReferenceNo, req.Status)
	if err != nil {
		s.log.Error("Failed to update transfer record status", err.Error())
		return model.Transaction{}, err
	}

	return transaction, nil
}

func (s *transactionService) ValidateTransferRecordCSV(ctx context.Context, records []entity.TransferRecordData) (res entity.ValidateTransferRecordCSVResponse, err error) {

	toAccountNos := make([]string, 0)
	totalTransferAmount := 0.0
	for _, record := range records {
		totalTransferAmount += record.Amount
		toAccountNos = append(toAccountNos, record.ToAccountNo)
	}

	users, err := s.userRepository.GetList(ctx, entity.GetListUserRequest{
		CorporateAccountNos: toAccountNos,
	})
	if err != nil {
		s.log.Error("Failed to get list of users", err.Error())
		return res, err
	}

	mapUserAccountNo := make(map[string]model.User)

	for _, user := range users {
		mapUserAccountNo[user.CorporateAccountNo] = user
	}

	errs := make([]string, 0)
	totalAccountNoFound := 0
	totalAccountNameNotMatch := 0
	wg := sync.WaitGroup{}
	wg.Add(len(records))
	for _, record := range records {
		go func(record entity.TransferRecordData) {
			defer wg.Done()
			user, ok := mapUserAccountNo[record.ToAccountNo]
			if !ok {
				totalAccountNoFound++
				return
			}

			if user.Name != record.ToAccountName {
				totalAccountNameNotMatch++
				return
			}
		}(record)
	}
	wg.Wait()

	if totalAccountNoFound > 0 {
		errs = append(errs, fmt.Sprintf("%d records error not found", totalAccountNoFound))
	}

	if totalAccountNameNotMatch > 0 {
		errs = append(errs, fmt.Sprintf("%d records error not match account name query by issuing bank", totalAccountNameNotMatch))
	}

	res = entity.ValidateTransferRecordCSVResponse{
		TransferRecords:     records,
		TotalTranferRecord:  len(records),
		TotalTransferAmount: totalTransferAmount,
		Errors:              errs,
	}

	return res, nil
}

func (s *transactionService) generateReferenceNo() string {
	rand.Seed(time.Now().UnixNano())

	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, 32)
	for i := range b {
		rand.Seed(time.Now().UnixNano())
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
