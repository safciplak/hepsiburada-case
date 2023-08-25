package orders

import (
	"context"
	"errors"

	"github.com/hepsiburada/hepsiburada-command/app"
	"gorm.io/gorm"
)

func NewMysqlRepo(db *gorm.DB) *MysqlRepo {
	return &MysqlRepo{db: db}
}

type MysqlRepo struct {
	db *gorm.DB
}

func (r *MysqlRepo) CreateOrder(ctx context.Context, o *Order) error {
	if err := r.db.Create(&o).Error; err != nil {
		return app.WrapError(err)
	}
	return nil
}

func (r *MysqlRepo) GetOrder(ctx context.Context, orderName string) (*Order, error) {
	var o Order
	if err := r.db.Where("name = ?", orderName).First(&o).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &o, ErrorNotfound
		}
		return &o, app.WrapError(err)
	}

	return &o, nil
}
