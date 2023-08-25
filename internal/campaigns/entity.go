package campaigns

import "time"

type Campaign struct {
	ID                     int64      `gorm:"column:id"`
	Name                   string     `gorm:"column:name"`
	ProductCode            string     `gorm:"column:product_code"`
	Duration               int64      `gorm:"column:duration"`
	PriceManipulationLimit int64      `gorm:"column:price_manipulation_limit"`
	TargetSalesCount       int64      `gorm:"column:target_sales_count"`
	CreatedAt              *time.Time `gorm:"column:created_at"`
	UpdatedAt              *time.Time `gorm:"column:updated_at"`
}

func (c Campaign) TableName() string {
	return "campaigns"
}
