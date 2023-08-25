package orders

import (
	"context"

	"github.com/hepsiburada/hepsiburada-command/app"
)

var (
	ErrorNotfound = app.BusinessError("order not found")
)

type orders interface {
	CreateOrder(ctx context.Context, o *Order) error
	GetOrder(ctx context.Context, orderName string) (*Order, error)
}

type Service struct {
	Orders orders
}

func (s *Service) CreateOrder(ctx context.Context, o *Order) error {
	return s.Orders.CreateOrder(ctx, o)
}

func (s *Service) GetOrder(ctx context.Context, orderName string) (*Order, error) {
	return s.Orders.GetOrder(ctx, orderName)
}
