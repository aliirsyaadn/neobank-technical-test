package repository

import (
	"context"

	"github.com/aliirsyaadn/neobank-technical-test/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type transferRecordRepository struct {
	log *zap.SugaredLogger
	db  *gorm.DB
}

func NewTransferRecordRepository(
	log *zap.SugaredLogger,
	db *gorm.DB,
) TransferRecordRepository {
	return &transferRecordRepository{
		log: log,
		db:  db,
	}
}

func (r *transferRecordRepository) GetList(ctx context.Context, referenceNo string) ([]model.TransferRecord, error) {
	var transferRecords []model.TransferRecord
	err := r.db.WithContext(ctx).Where("transaction_reference_no = ?", referenceNo).Find(&transferRecords).Error
	if err != nil {
		r.log.Error("Failed to get list of transfer records", err.Error())
		return nil, err
	}
	return transferRecords, nil
}

func (r *transferRecordRepository) InsertBatch(ctx context.Context, transferRecords *[]model.TransferRecord) error {
	err := r.db.WithContext(ctx).Create(transferRecords).Error
	if err != nil {
		r.log.Error("Failed to insert transfer records", err.Error())
		return err
	}
	return nil
}

func (r *transferRecordRepository) UpdateStatus(ctx context.Context, referenceNo string, status string) error {
	err := r.db.WithContext(ctx).Model(&model.TransferRecord{}).Where("transaction_reference_no = ?", referenceNo).Update("status", status).Error
	if err != nil {
		r.log.Error("Failed to update transfer record status", err.Error())
		return err
	}
	return nil
}
