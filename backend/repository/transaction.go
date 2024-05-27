package repository

import (
	"context"

	"github.com/aliirsyaadn/neobank-technical-test/entity"
	"github.com/aliirsyaadn/neobank-technical-test/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type transactionRepository struct {
	log *zap.SugaredLogger
	db  *gorm.DB
}

func NewTransactionRepository(
	log *zap.SugaredLogger,
	db *gorm.DB,
) TransactionRepository {
	return &transactionRepository{
		log: log,
		db:  db,
	}
}

func (r *transactionRepository) Get(ctx context.Context, referenceNo string) (model.Transaction, error) {
	var transaction model.Transaction
	err := r.db.WithContext(ctx).Where("reference_no = ?", referenceNo).Preload("MakerUserDetail").First(&transaction).Error
	if err != nil {
		r.log.Error("Failed to get transaction", err.Error())
		return model.Transaction{}, err
	}
	return transaction, nil
}

func (r *transactionRepository) GetList(ctx context.Context, req entity.GetListTransactionRequest) (res entity.GetListTransactionResponse, err error) {
	query := r.db.WithContext(ctx)

	if req.Status != "" {
		query = query.Where("status = ?", req.Status)
	}

	if req.FromAccountNo != "" {
		query = query.Where("from_account_no = ?", req.FromAccountNo)
	}

	if !req.StartDate.IsZero() && !req.EndDate.IsZero() {
		query = query.Where("created_at BETWEEN ? AND ?", req.StartDate, req.EndDate)
	}

	var transactions []model.Transaction
	query = query.
		Scopes(PaginateQuery(&req.Pagination)).
		Find(&transactions)

	err = query.Error
	if err != nil {
		r.log.Error("Failed to get list of transactions", err.Error())
		return res, err
	}

	var total int64
	err = query.Model(&model.Transaction{}).Count(&total).Error
	if err != nil {
		r.log.Error("Failed to get total of transactions", err.Error())
		return res, err
	}

	req.Pagination.DataPerPage = len(transactions)
	req.Pagination.TotalData = int(total)

	res = entity.GetListTransactionResponse{
		Transactions: transactions,
		Pagination:   req.Pagination,
	}

	return res, nil
}

func (r *transactionRepository) Insert(ctx context.Context, transaction *model.Transaction) error {
	err := r.db.WithContext(ctx).Create(transaction).Error
	if err != nil {
		r.log.Error("Failed to insert transaction", err.Error())
		return err
	}
	return nil
}

func (r *transactionRepository) Update(ctx context.Context, transaction *model.Transaction) error {
	err := r.db.WithContext(ctx).Omit("CreatedAt").Updates(transaction).Error
	if err != nil {
		r.log.Error("Failed to update transaction", err.Error())
		return err
	}
	return nil
}
