package repository

import (
	"context"

	"github.com/aliirsyaadn/neobank-technical-test/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type userOTPRepository struct {
	log *zap.SugaredLogger
	db  *gorm.DB
}

func NewUserOTPRepository(
	log *zap.SugaredLogger,
	db *gorm.DB,
) UserOTPRepository {
	return &userOTPRepository{
		log: log,
		db:  db,
	}
}

func (r *userOTPRepository) Get(ctx context.Context, email, otp string) (model.UserOTP, error) {
	var otpUser model.UserOTP
	err := r.db.WithContext(ctx).Where("email = ? AND otp = ?", email, otp).First(&otpUser).Error
	if err != nil {
		r.log.Error("Failed to get otp user", err.Error())
		return model.UserOTP{}, err
	}
	return otpUser, nil
}

func (r *userOTPRepository) Insert(ctx context.Context, otpUser *model.UserOTP) error {
	err := r.db.WithContext(ctx).Create(otpUser).Error
	if err != nil {
		r.log.Error("Failed to insert otp user", err.Error())
		return err
	}
	return nil
}
