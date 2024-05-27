package repository

import (
	"context"

	"github.com/aliirsyaadn/neobank-technical-test/entity"
	"github.com/aliirsyaadn/neobank-technical-test/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type userRepository struct {
	log *zap.SugaredLogger
	db  *gorm.DB
}

func NewUserRepository(
	log *zap.SugaredLogger,
	db *gorm.DB,
) UserRepository {
	return &userRepository{
		log: log,
		db:  db,
	}
}

func (r *userRepository) GetList(ctx context.Context, req entity.GetListUserRequest) ([]model.User, error) {
	var users []model.User
	query := r.db.WithContext(ctx)

	if len(req.CorporateAccountNos) > 0 {
		query = query.Where("corporate_account_no IN ?", req.CorporateAccountNos)
	}

	err := query.Find(&users).Error
	if err != nil {
		r.log.Error("Failed to get list of users", err.Error())
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Get(ctx context.Context, id string) (model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		r.log.Error("Failed to get user", err.Error())
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		r.log.Error("Failed to get user", err.Error())
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) Insert(ctx context.Context, user *model.User) error {
	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		r.log.Error("Failed to insert user", err.Error())
		return err
	}
	return nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	err := r.db.WithContext(ctx).Omit("CreatedAt").Updates(user).Error
	if err != nil {
		r.log.Error("Failed to update user", err.Error())
		return err
	}
	return nil
}
