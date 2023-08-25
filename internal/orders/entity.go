package orders

import "time"

type Order struct {
	ID          int64      `gorm:"column:id"`
	ProductCode string     `gorm:"column:product_code"`
	Quantity    int64      `gorm:"column:quantity"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
}

func (o Order) TableName() string {
	return "orders"
}
