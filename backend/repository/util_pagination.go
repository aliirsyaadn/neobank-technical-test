package repository

import (
	"github.com/aliirsyaadn/neobank-technical-test/entity"
	"gorm.io/gorm"
)

func PaginateQuery(p *entity.Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if p.Page == -1 {
			return db
		}

		if p.Page == 0 {
			p.Page = 1
		}

		if p.PageSize > 10 || p.PageSize <= 0 {
			p.PageSize = 10
		}

		offset := (p.Page - 1) * p.PageSize
		return db.Offset(offset).Limit(p.PageSize)
	}
}
