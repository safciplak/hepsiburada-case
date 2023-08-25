package products

import "time"

type Product struct {
	ID          int64      `gorm:"column:id"`
	ProductCode string     `gorm:"column:product_code"`
	Price       int64      `gorm:"column:price"`
	Stock       int64      `gorm:"column:stock"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
}

func (p Product) TableName() string {
	return "products"
}
