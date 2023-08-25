package campaigns

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

func (r *MysqlRepo) CreateCampaign(ctx context.Context, c *Campaign) error {
	if err := r.db.Create(&c).Error; err != nil {
		return app.WrapError(err)
	}
	return nil
}

func (r *MysqlRepo) GetCampaign(ctx context.Context, campaignName string) (*Campaign, error) {
	var c Campaign
	if err := r.db.Where("name = ?", campaignName).First(&c).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &c, ErrorNotfound
		}
		return &c, app.WrapError(err)
	}

	return &c, nil
}
