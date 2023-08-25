package products

import (
	"github.com/hepsiburada/hepsiburada-command/app"
	"context"
	"errors"
	"gorm.io/gorm"
)

func NewMysqlRepo(db *gorm.DB) *MysqlRepo {
	return &MysqlRepo{db: db}
}

type MysqlRepo struct {
	db *gorm.DB
}

func (r *MysqlRepo) CreateProduct(ctx context.Context, p *Product) error {
	if err := r.db.Create(p).Error; err != nil {
		return app.WrapError(err)
	}
	return nil
}

func (r *MysqlRepo) GetProduct(ctx context.Context, productCode string) (*Product, error) {
	var p Product
	if err := r.db.Where("product_code = ?", productCode).First(&p).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &p, ErrorNotfound
		}
		return &p, app.WrapError(err)
	}

	return &p, nil
}
